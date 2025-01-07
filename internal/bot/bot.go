package bot

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/your_project/internal/handlers"
)

func StartBot() {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	// 注册二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 使用消息分发器
	dispatcher := openwechat.NewMessageMatchDispatcher()

	// 注册各类消息处理器
	dispatcher.OnText(handlers.PingHandler)
	dispatcher.OnText(handlers.MusicHandler)
	dispatcher.OnText(handlers.TextHandler)

	// 将分发器注册为消息处理器
	bot.MessageHandler = dispatcher.AsMessageHandler()

	// 登录
	if err := bot.Login(); err != nil {
		fmt.Println("登录失败:", err)
		return
	}

	// 阻塞主goroutine, 直到手动退出
	bot.Block()
}
