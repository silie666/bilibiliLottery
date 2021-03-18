package main

import (
	"bilibili/function"
	"time"
)

func main() {
	//function.BilibiliGetDo() //更新抽奖列表
	for true {
		time.Sleep(24*time.Hour)
		function.BilibliDoUpdate()
		function.BilibiliDoRun()
	}
}