//go:build wireinject
// +build wireinject

/*
 * @FilePath: /sxwz-bot/cmd/wire.go
 * @Author: maggot-code
 * @Date: 2023-09-18 09:31:05
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 04:11:33
 * @Description:
 */
package main

import (
	"github.com/google/wire"
	"github.com/maggot-code/sxwz-bot/internal/conf"
	"github.com/maggot-code/sxwz-bot/internal/gateway"
	"github.com/maggot-code/sxwz-bot/internal/handler"
)

func wireApp(*conf.Bootstrap) (*App, func(), error) {
	panic(wire.Build(
		gateway.ProviderSet,
		handler.ProviderSet,
		new,
	))
}
