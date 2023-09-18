/*
 * @FilePath: /sxwz-bot/internal/biz/biz.go
 * @Author: maggot-code
 * @Date: 2023-09-18 12:58:41
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 13:02:10
 * @Description:
 */
package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewATMessageUseCase)
