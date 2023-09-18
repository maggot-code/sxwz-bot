package handler

import (
	"github.com/maggot-code/sxwz-bot/internal/biz"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
)

type ATMessageEventHandler struct {
	ac *biz.ATMessageUseCase
}

func NewATMessageEventHandler(ac *biz.ATMessageUseCase) event.ATMessageEventHandler {
	ah := &ATMessageEventHandler{
		ac: ac,
	}

	return ah.Handle
}

func (ah *ATMessageEventHandler) Handle(event *dto.WSPayload, data *dto.WSATMessageData) error {
	return ah.ac.Reply(data)
}
