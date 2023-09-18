/*
 * @FilePath: /sxwz-bot/cmd/main.go
 * @Author: maggot-code
 * @Date: 2023-09-18 08:17:19
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 09:07:11
 * @Description:
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"syscall"
	"time"

	"github.com/maggot-code/sxwz-bot/internal/conf"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	conf, err := conf.New(flagconf)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	botToken := token.BotToken(uint64(conf.Bot.Appid), conf.Bot.Token)
	api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)

	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln(err)
	}

	websocket.RegisterResumeSignal(syscall.SIGUSR1)
	intent := websocket.RegisterHandlers(
		// 连接关闭回调
		ErrorNotifyHandler(),
		// 如果想要捕获到连接成功的事件，可以实现这个回调
		ReadyHandler(),
		// 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		CreateMessageHandler(),
		// at 机器人事件，目前是在这个事件处理中有逻辑，会回消息，其他的回调处理都只把数据打印出来，不做任何处理
		ATMessageEventHandler(),
	)

	// 指定需要启动的分片数为 2 的话可以手动修改 wsInfo
	if err = botgo.NewSessionManager().Start(ws, botToken, &intent); err != nil {
		log.Fatalln(err)
	}
}

func ErrorNotifyHandler() event.ErrorNotifyHandler {
	return func(err error) {
		log.Println("error notify receive: ", err)
	}
}

func ReadyHandler() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		log.Println("ready event receive: ", data)
	}
}

// CreateMessageHandler 处理消息事件
func CreateMessageHandler() event.MessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSMessageData) error {
		fmt.Println(data)
		return nil
	}
}

// ATMessageEventHandler 实现处理 at 消息的回调
func ATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		input := strings.ToLower(message.ETLInput(data.Content))
		fmt.Println(input)

		return nil
	}
}
