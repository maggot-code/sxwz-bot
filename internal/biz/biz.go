/*
 * @FilePath: /sxwz-bot/internal/biz/biz.go
 * @Author: maggot-code
 * @Date: 2023-09-19 04:47:36
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 04:50:09
 * @Description:
 */
package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMessageUseCase)
