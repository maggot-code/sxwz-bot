/*
 * @FilePath: /sxwz-bot/cmd/main.go
 * @Author: maggot-code
 * @Date: 2023-09-18 08:17:19
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 10:38:38
 * @Description:
 */
package main

import (
	"flag"

	"github.com/maggot-code/sxwz-bot/internal/conf"
	"github.com/maggot-code/sxwz-bot/internal/gateway"
	"github.com/tencent-connect/botgo/dto"
)

var flagconf string

type App struct {
	gateway *gateway.Gateway
	intent  dto.Intent
}

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func new(gw *gateway.Gateway, intent dto.Intent) *App {
	return &App{gateway: gw, intent: intent}
}

func main() {
	flag.Parse()

	conf, err := conf.New(flagconf)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(conf)
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err := app.gateway.Start(app.intent); err != nil {
		panic(err)
	}
}
