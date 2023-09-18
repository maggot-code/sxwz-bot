/*
 * @FilePath: /sxwz-bot/internal/handler/handler.go
 * @Author: maggot-code
 * @Date: 2023-09-19 02:45:02
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 04:30:59
 * @Description:
 */
package handler

import (
	"reflect"

	"github.com/google/wire"
	"github.com/tencent-connect/botgo/event"
)

type Handler struct {
	Ready       event.ReadyHandler
	ErrorNotify event.ErrorNotifyHandler

	ATMessage event.ATMessageEventHandler
}

var ProviderSet = wire.NewSet(
	NewReadyHandler,
	NewErrorNotifyHandler,
	NewATMessageEventHandler,
	wire.Struct(new(Handler), "*"),
	NewHandler)

func NewHandler(h Handler) []interface{} {
	// 使用反射获取结构体字段值并构建切片
	var handlers []interface{}
	val := reflect.ValueOf(h)

	for i := 0; i < val.NumField(); i++ {
		handlers = append(handlers, val.Field(i).Interface())
	}

	return handlers
}
