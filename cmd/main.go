package main

import (
	"fmt"
	"wechatBot/internal/bot"
)

func main() {
	// 启动微信Bot
	fmt.Println("正在启动微信Bot...........")
	bot.StartBot()
}
