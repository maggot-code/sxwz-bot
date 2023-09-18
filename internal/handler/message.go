/*
 * @FilePath: /sxwz-bot/internal/handler/message.go
 * @Author: maggot-code
 * @Date: 2023-09-19 04:21:22
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 04:22:12
 * @Description:
 */
package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/log"
)

func NewATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		log.Info("at message handler")
		return nil
	}
}
