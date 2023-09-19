/*
 * @FilePath: /sxwz-bot/internal/data/data.go
 * @Author: maggot-code
 * @Date: 2023-09-19 04:51:51
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 07:18:32
 * @Description:
 */
package data

import (
	"github.com/google/wire"
	"github.com/maggot-code/sxwz-bot/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewCommandStore, NewMessageRepo)

type Data struct{}

func NewData(c *conf.Bootstrap) (*Data, func(), error) {
	return &Data{}, func() {}, nil
}
