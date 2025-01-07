package handlers

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

func PingHandler(ctx *openwechat.MessageContext) {
	msg := ctx.Message
	if msg.IsText() && msg.Content == "ping" {
		fmt.Println("收到Ping消息")
		msg.ReplyText("pong")
	}
}
