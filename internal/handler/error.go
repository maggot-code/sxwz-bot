package handler

import (
	"log"

	"github.com/maggot-code/sxwz-bot/internal/gateway"
	"github.com/tencent-connect/botgo/event"
)

func NewErrorNotifyHandler(gw *gateway.Gateway) event.ErrorNotifyHandler {
	return func(err error) {
		log.Println(err)
	}
}
