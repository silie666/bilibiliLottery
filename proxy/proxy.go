package proxy

import (
	"encoding/json"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
	"bilibili/common"
	"bilibili/drivers/redis"
	"bilibili/logger"
)

type Proxy struct {
	Code int64 `json:"code"`
	Data []struct{
		Ip string `json:"ip"`
		Port int64 `json:"port"`
		ExpireTime string `json:"expire_time"`
	}
}

var proxyUrl string = "http://webapi.http.zhimacangku.com/getip?num=1&type=2&pro=&city=0&yys=0&port=1&pack=134598&ts=1&ys=0&cs=0&lb=1&sb=0&pb=4&mr=1&regions="

//定时更新
func ProxyTicker() {
	var wg sync.WaitGroup
	wg.Add(1)
	ticker := time.NewTicker(5 * time.Second)
	i := 0
	go func(t *time.Ticker) {
		defer wg.Done()
		for  {
			<-ticker.C
			i++
			data,err := GetProxyIp()
			if err != nil {
				logger.LoggerToFile(err.Error())
				ticker.Stop()
				return
			}
			key := "proxy_ip_" + data["ip"].(string) + ":" + strconv.FormatInt(data["port"].(int64), 10)
			now := time.Now().Unix()
			expireTime := common.TimeStamp(data["expire_time"].(string))
			ex := expireTime - now
			_,err = redis.RedisDb.Do("HMSet",redigo.Args{}.Add(key).AddFlat(data)...)
			redis.RedisDb.Do("EXPIRE",key,ex)
			if err != nil {
				logger.LoggerToFile("Redis Err:"+err.Error())
				ticker.Stop()
				return
			}
			if i == 20 {//免费版本一天20个
				ticker.Stop()
				return
			}
			fmt.Println("定时获取成功")
		}
	}(ticker)
	wg.Wait()
}

//通过url获取最新
func GetProxyIp() (map[string]interface{},error) {
	var proxy Proxy
	resp,_ := http.Get(proxyUrl)
	response,_ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(response,&proxy)

	data := make(map[string]interface{},0)
	if len(proxy.Data) == 0 {
		return nil,fmt.Errorf("获取代理ip失败")
	}else{
		data["ip"] = proxy.Data[0].Ip
		data["port"] = proxy.Data[0].Port
		data["expire_time"] = proxy.Data[0].ExpireTime
	}
	return data,nil
}

//通过redis获取
func GetProxyRedis(_ *http.Request)(*url.URL,error)  {
	res,_ := redigo.Strings(redis.RedisDb.Do("KEYS","proxy_ip_*"))
	if len(res) == 0 {
		return nil,fmt.Errorf("无ip获取")
	}
	var proxyURLs = make([]*url.URL, 0)
	for _, v := range res {
		proxy, _ := redigo.StringMap(redis.RedisDb.Do("HGetAll",v))
		proxyURLs = append(proxyURLs, &url.URL{
			Host: proxy["ip"] + ":" + proxy["port"],
		})
	}
	fmt.Println(proxyURLs)
	rand.Seed(time.Now().Unix())
	return proxyURLs[rand.Intn(len(proxyURLs))], nil
}

