/*
 * @FilePath: /sxwz-bot/internal/biz/message.go
 * @Author: maggot-code
 * @Date: 2023-09-19 04:47:49
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 06:38:14
 * @Description:
 */
package biz

import (
	"github.com/tencent-connect/botgo/dto"
)

type MessageRepo interface {
	ReplyATMessage(data *dto.WSATMessageData) error
}

type MessageUseCase struct {
	repo      MessageRepo
	atmessage chan *dto.WSATMessageData
}

func NewMessageUseCase(mr MessageRepo) *MessageUseCase {
	return &MessageUseCase{
		repo:      mr,
		atmessage: make(chan *dto.WSATMessageData, 128),
	}
}

func (mc *MessageUseCase) LoopATMessage() {
	for data := range mc.atmessage {
		go mc.repo.ReplyATMessage(data)
	}
}

func (mc *MessageUseCase) PushATMessage(data *dto.WSATMessageData) {
	mc.atmessage <- data
}
