/*
 * @FilePath: /sxwz-bot/internal/biz/at.go
 * @Author: maggot-code
 * @Date: 2023-09-18 13:01:20
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 14:19:41
 * @Description:
 */
package biz

import (
	"context"
	"strings"

	"github.com/maggot-code/sxwz-bot/internal/gateway"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

type ATMessageUseCase struct {
	gw   *gateway.Gateway
	repo ATMessageRepo
}

type ATMessageRepo interface {
	CareerData(data *dto.WSATMessageData) (*dto.MessageToCreate, error)
}

func NewATMessageUseCase(gw *gateway.Gateway, ar ATMessageRepo) *ATMessageUseCase {
	return &ATMessageUseCase{gw: gw, repo: ar}
}

func (ac *ATMessageUseCase) Reply(data *dto.WSATMessageData) error {
	ctx := context.Background()
	input := strings.ToLower(message.ETLInput(data.Content))
	cmd := message.ParseCommand(input)

	mapping := map[string]func(data *dto.WSATMessageData) (*dto.MessageToCreate, error){
		"/事业民": ac.repo.CareerData,
	}

	if fn, ok := mapping[cmd.Cmd]; ok {
		msg, err := fn(data)
		if err != nil {
			return err
		}

		ac.gw.Send(ctx, data.ChannelID, msg)
	}

	return nil
}
