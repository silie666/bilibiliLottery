package common

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
	"strings"
	"time"
)



func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122  {
		strArry[0] -=  32
	}
	return string(strArry)
}

func JsonEncode(data interface{})string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return string(buffer.Bytes())
}

func JsonDecode(json []byte)*simplejson.Json {
	js,_ := simplejson.NewJson(json)
	return js
}


func TimeStamp(timeString string)int64 {
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, timeString, loc)
	timestamp := tmp.Unix()    //转化为时间戳 类型是int64
	return timestamp
}

func JsonAtoi(str interface{}) int {
	str = str.(json.Number).String()
	str,_ = strconv.Atoi(str.(string))
	return str.(int)
}

func LSJPrice(str string) float64 {
	if str == "暂无" || str == "-" {
		return 0.00
	}
	strIndex:= strings.Index(str,"万")
	num,_ := strconv.ParseFloat(str[:strIndex],10)
	return num
}


func RsaEncrypt(origData,publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, fmt.Errorf("加密失败")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}