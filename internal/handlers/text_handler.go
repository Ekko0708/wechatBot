package handlers

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

// 普通文本消息处理器
func TextHandler(ctx *openwechat.MessageContext) {
	msg := ctx.Message
	fmt.Printf("收到文本消息: %s\n", msg.Content)
	msg.ReplyText("收到你的消息！")
}
