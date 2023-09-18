/*
 * @FilePath: /sxwz-bot/internal/gateway/gateway.go
 * @Author: maggot-code
 * @Date: 2023-09-18 10:16:19
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 14:19:36
 * @Description:
 */
package gateway

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/maggot-code/sxwz-bot/internal/conf"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
)

var ProviderSet = wire.NewSet(NewGateway)

type Gateway struct {
	api   openapi.OpenAPI
	token *token.Token
	ws    *dto.WebsocketAP
}

func NewGateway(c *conf.Bootstrap) *Gateway {
	ctx := context.Background()
	botToken := token.BotToken(uint64(c.Bot.Appid), c.Bot.Token)
	api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)

	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		panic(err)
	}

	return &Gateway{
		api:   api,
		token: botToken,
		ws:    ws,
	}
}

func (gw *Gateway) Start(intent dto.Intent) error {
	// intent := websocket.RegisterHandlers()

	return botgo.NewSessionManager().Start(gw.ws, gw.token, &intent)
}

func (gw *Gateway) Send(ctx context.Context, channelID string, msg *dto.MessageToCreate) error {
	_, err := gw.api.PostMessage(ctx, channelID, msg)
	if err != nil {
		return err
	}

	return nil
}
