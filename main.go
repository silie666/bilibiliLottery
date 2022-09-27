package main

import (
	"bilibili/function"
	"fmt"
	"os"
	"time"
)

func main() {

	osArr := os.Args

	switch osArr[1] {
	case "-init":
		function.Init()
		break
	case "-login":
		function.BilibiliLogin()
		break
	case "-draw":
		function.BilibiliGeTLuckDraw()
		function.BilibliDoUpdate()
		function.BilibiliDoRun()
		break
	case "-del":
		for true {
			function.BilibiliAutoDel()
			time.Sleep(30 * time.Second)
		}
		break
	case "-start":
		function.BilibiliStart()
		break
	case "-cancel-modify":
		function.BilibiliCancelModify()
	default:
		fmt.Println("请输入命令，-draw开始抽奖，-del删除动态，-forward开始转发动态抽奖")
	}

}
