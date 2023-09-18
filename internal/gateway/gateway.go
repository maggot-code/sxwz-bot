/*
 * @FilePath: /sxwz-bot/internal/gateway/gateway.go
 * @Author: maggot-code
 * @Date: 2023-09-18 10:16:19
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 02:43:23
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
	ctx   context.Context
	api   openapi.OpenAPI
	token *token.Token
	ws    *dto.WebsocketAP
}

func NewGateway(c *conf.Bootstrap) *Gateway {
	botToken := token.BotToken(uint64(c.Bot.Appid), c.Bot.Token)
	api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()

	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		panic(err)
	}

	return &Gateway{
		ctx:   ctx,
		api:   api,
		token: botToken,
		ws:    ws,
	}
}

func (gw *Gateway) Start(intent dto.Intent) error {
	return botgo.NewSessionManager().Start(gw.ws, gw.token, &intent)
}
