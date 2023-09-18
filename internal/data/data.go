package data

import (
	"github.com/google/wire"
	"github.com/maggot-code/sxwz-bot/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewATMessageRepo)

type Data struct {
	conf *conf.Bootstrap
}

func NewData(c *conf.Bootstrap) (*Data, func(), error) {
	return &Data{conf: c}, func() {}, nil
}
