package data

import (
	"fmt"
	"sync"

	"github.com/maggot-code/sxwz-bot/internal/conf"
	"github.com/maggot-code/sxwz-bot/internal/gateway"
	"github.com/maggot-code/sxwz-bot/pkg/fans"
	"github.com/tencent-connect/botgo/dto"
)

type CommandStore struct {
	conf *conf.Bootstrap
	gw   *gateway.Gateway
}

func NewCommandStore(c *conf.Bootstrap, gw *gateway.Gateway) *CommandStore {
	return &CommandStore{conf: c, gw: gw}
}

func (cs *CommandStore) NoneCommand(data *dto.WSATMessageData) error {
	return cs.gw.PostMessage(data.ChannelID, &dto.MessageToCreate{
		Content: "啊？🤔️",
		MsgID:   data.ID,
		MessageReference: &dto.MessageReference{
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	})
}

type CareerContent struct {
	total string
	real  string
}

func (cs *CommandStore) CareerData(data *dto.WSATMessageData) error {
	var wg sync.WaitGroup
	cache := make(chan CareerContent, len(cs.conf.Observe))

	for _, uid := range cs.conf.Observe {
		wg.Add(1)
		go func(id int32) {
			defer wg.Done()
			f, err := fans.NewFansRepo(id)
			if err != nil {
				fmt.Println(err)
			}

			cache <- CareerContent{
				total: fmt.Sprintf("%s\n", f.ToUserTotal()),
				real:  fmt.Sprintf("%s\n", f.ToUserTrue()),
			}
		}(uid)
	}

	wg.Wait()
	close(cache)

	var total, real string
	for item := range cache {
		total += item.total
		real += item.real
	}

	return cs.gw.PostMessage(data.ChannelID, &dto.MessageToCreate{
		Content: fmt.Sprintf("\n%s", total),
		MsgID:   data.ID,
		MessageReference: &dto.MessageReference{
			MessageID:             data.ID,
			IgnoreGetMessageError: false,
		},
	})
}
