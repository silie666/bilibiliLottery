package common

import (
	"bilibili/config"
	"bilibili/logger"
	"bilibili/respdata"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

func JsonEncode(data interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return string(buffer.Bytes())
}

func JsonDecode(json []byte) *simplejson.Json {
	js, _ := simplejson.NewJson(json)
	return js
}

func TimeStamp(timeString string) int64 {
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, timeString, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	return timestamp
}

func CurrentTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func JsonAtoi(str interface{}) int {
	str = str.(json.Number).String()
	str, _ = strconv.Atoi(str.(string))
	return str.(int)
}

func Get(apiUrl string, data url.Values, cookie ...http.Cookie) (response string) {
	parseURL, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}
	if data != nil {
		parseURL.RawQuery = data.Encode()
	}
	client := &http.Client{Timeout: 5 * time.Second}
	var req *http.Request
	req, _ = http.NewRequest("GET", parseURL.String(), nil)
	cookie1 := &http.Cookie{Name: "SESSDATA", Value: config.Env.GetString("cookie.sess_data")}
	req.AddCookie(cookie1)
	for _, v := range cookie {
		req.AddCookie(&v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	resp, error := client.Do(req)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	response = result.String()
	return
}

func Post(apiUrl string, data interface{}, contentType string, cookie ...http.Cookie) (response string) {
	var req *http.Request
	var err error
	value, ok := data.(url.Values)
	if ok {
		req, err = http.NewRequest("POST", apiUrl, strings.NewReader(value.Encode()))
	} else {
		body, _ := json.Marshal(data)
		req, err = http.NewRequest("POST", apiUrl, strings.NewReader(string(body)))
	}
	req.Header.Set("content-type", contentType)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	cookie1 := &http.Cookie{Name: "SESSDATA", Value: config.Env.GetString("cookie.sess_data")}
	req.AddCookie(cookie1)
	for _, v := range cookie {
		req.AddCookie(&v)
	}
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	response = result.String()
	return
}

func ShortUrlRedirect(shortUrl string) string {
	redirectCount := 0
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) (e error) {
		redirectCount++
		if redirectCount == 1 {
			return errors.New(req.URL.String())
		}
		return
	}}
	response, err := client.Get(shortUrl)
	if err != nil {
		if e, ok := err.(*url.Error); ok && e.Err != nil {
			remoteUrl := e.URL
			return remoteUrl
		}
	}
	defer response.Body.Close()
	return shortUrl
}

func BilibiliIsError(response string) error {
	var bilibiliCode respdata.BilibiliCode
	json.Unmarshal([]byte(response), &bilibiliCode)
	if bilibiliCode.Code != 0 {
		logger.LogToFile("错误：" + bilibiliCode.Message)
		return fmt.Errorf("错误：" + bilibiliCode.Message)
	}
	return nil
}
