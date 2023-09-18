/*
 * @FilePath: /sxwz-bot/internal/handler/use.go
 * @Author: maggot-code
 * @Date: 2023-09-19 04:15:41
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 04:31:06
 * @Description:
 */
package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/log"
)

func NewReadyHandler() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		log.Info("ready handler")
	}
}

func NewErrorNotifyHandler() event.ErrorNotifyHandler {
	return func(err error) {
		log.Error(err)
	}
}
