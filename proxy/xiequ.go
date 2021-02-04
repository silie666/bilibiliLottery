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
	"bilibili/drivers/redis"
	"bilibili/logger"
)

type XiequProxy struct {
	Code int64 `json:"code"`
	Data []struct{
		Ip string `json:"IP"`
		Port int64 `json:"Port"`
	}
}

var XiequProxyUrl string = "http://api.xiequ.cn/VAD/GetIp.aspx?act=get&num=200&time=30&plat=0&re=1&type=0&so=1&ow=1&spl=1&addr=&db=1"

//定时更新
func XiequTicker() {
	var wg sync.WaitGroup
	wg.Add(1)
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	go func(t *time.Ticker) {
		defer wg.Done()
		for  {
			<-ticker.C
			i++
			data,err := GetXiequIp()
			if err != nil {
				logger.LoggerToFile(err.Error())
				ticker.Stop()
				return
			}
			key := "xiequ_proxy_ip_" + data["ip"].(string) + ":" + strconv.FormatInt(data["port"].(int64), 10)
			_,err = redis.RedisDb.Do("HMSet",redigo.Args{}.Add(key).AddFlat(data)...)
			redis.RedisDb.Do("EXPIRE",key,30)
			if err != nil {
				logger.LoggerToFile("Redis Err:"+err.Error())
				ticker.Stop()
				return
			}
			if i == 3 {//免费版本一天20个
				ticker.Stop()
				return
			}
			fmt.Println("定时获取成功")
		}
	}(ticker)
	wg.Wait()
}

//通过url获取最新
func GetXiequIp() (map[string]interface{},error) {
	var proxy XiequProxy
	resp,_ := http.Get(XiequProxyUrl)
	response,_ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(response,&proxy)

	data := make(map[string]interface{},0)
	if len(proxy.Data) == 0 {
		return nil,fmt.Errorf("获取代理ip失败")
	}else{
		data["ip"] = proxy.Data[0].Ip
		data["port"] = proxy.Data[0].Port
	}
	return data,nil
}

//通过redis获取 _ *http.Request
func GetXiequProxyRedis(_ *http.Request)(*url.URL,error)  {
	res,_ := redigo.Strings(redis.RedisDb.Do("KEYS","xiequ_proxy_ip_*"))
	fmt.Println(res)
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

