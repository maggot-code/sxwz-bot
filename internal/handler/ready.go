/*
 * @FilePath: /sxwz-bot/internal/handler/ready.go
 * @Author: maggot-code
 * @Date: 2023-09-18 12:50:51
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 12:51:09
 * @Description:
 */
package handler

import (
	"github.com/maggot-code/sxwz-bot/internal/gateway"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
)

func NewReadyHandler(gw *gateway.Gateway) event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		// do something
	}
}
