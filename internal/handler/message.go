/*
 * @FilePath: /sxwz-bot/internal/handler/message.go
 * @Author: maggot-code
 * @Date: 2023-09-19 04:21:22
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 06:37:09
 * @Description:
 */
package handler

import (
	"github.com/maggot-code/sxwz-bot/internal/biz"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
)

type MessageEventHandler struct {
	mc *biz.MessageUseCase
}

func NewATMessageEventHandler(mc *biz.MessageUseCase) event.ATMessageEventHandler {
	mh := &MessageEventHandler{mc: mc}

	go mh.mc.LoopATMessage()

	return mh.bind
}

func (mh *MessageEventHandler) bind(event *dto.WSPayload, data *dto.WSATMessageData) error {
	mh.mc.PushATMessage(data)
	return nil
}
