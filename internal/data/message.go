/*
 * @FilePath: /sxwz-bot/internal/data/message.go
 * @Author: maggot-code
 * @Date: 2023-09-19 04:53:14
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 07:30:27
 * @Description:
 */
package data

import (
	"strings"

	"github.com/maggot-code/sxwz-bot/internal/biz"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

var _ biz.MessageRepo = (*messageRepo)(nil)

type messageRepo struct {
	data *Data
	cs   *CommandStore
	// gw   *gateway.Gateway
}

func NewMessageRepo(data *Data, cs *CommandStore) biz.MessageRepo {
	return &messageRepo{data: data, cs: cs}
}

func (mr *messageRepo) ReplyATMessage(data *dto.WSATMessageData) error {
	command := strings.ToLower(message.ETLInput(data.Content))

	switch command {
	case "/事业民":
		return mr.cs.CareerData(data)
	default:
		return mr.cs.NoneCommand(data)
	}
}
