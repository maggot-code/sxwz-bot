//go:build wireinject
// +build wireinject

/*
 * @FilePath: /sxwz-bot/cmd/wire.go
 * @Author: maggot-code
 * @Date: 2023-09-18 09:31:05
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 13:47:44
 * @Description:
 */
package main

import (
	"github.com/google/wire"
	"github.com/maggot-code/sxwz-bot/internal/biz"
	"github.com/maggot-code/sxwz-bot/internal/conf"
	"github.com/maggot-code/sxwz-bot/internal/data"
	"github.com/maggot-code/sxwz-bot/internal/gateway"
	"github.com/maggot-code/sxwz-bot/internal/handler"
	"github.com/tencent-connect/botgo/websocket"
)

func wireApp(*conf.Bootstrap) (*App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		gateway.ProviderSet,
		handler.ProviderSet,
		websocket.RegisterHandlers,
		new,
	))
}
