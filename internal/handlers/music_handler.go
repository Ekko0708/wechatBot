package handlers

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

// 点歌处理器
func MusicHandler(ctx *openwechat.MessageContext) {
	msg := ctx.Message
	if strings.HasPrefix(msg.Content, "点歌") {
		songName := strings.TrimPrefix(msg.Content, "点歌 ")
		fmt.Printf("点歌请求: %s\n", songName)
		msg.ReplyText(fmt.Sprintf("正在搜索歌曲: %s", songName))
	}
}
