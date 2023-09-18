package handler

import (
	"github.com/google/wire"
	"github.com/tencent-connect/botgo/event"
)

var ProviderSet = wire.NewSet(
	NewErrorNotifyHandler,
	NewReadyHandler,
	NewATMessageEventHandler,
	NewHandler)

func NewHandler(
	eh event.ErrorNotifyHandler,
	rh event.ReadyHandler,
	ah event.ATMessageEventHandler) []interface{} {

	return []interface{}{
		eh,
		rh,
		ah,
	}
}
