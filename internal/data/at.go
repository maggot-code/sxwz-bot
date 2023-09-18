/*
 * @FilePath: /sxwz-bot/internal/data/at.go
 * @Author: maggot-code
 * @Date: 2023-09-18 13:43:30
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-18 14:20:38
 * @Description:
 */
package data

import (
	"github.com/maggot-code/sxwz-bot/internal/biz"
	"github.com/tencent-connect/botgo/dto"
)

var _ biz.ATMessageRepo = (*aTMessageRepo)(nil)

type aTMessageRepo struct {
	data *Data
}

func NewATMessageRepo(data *Data) biz.ATMessageRepo {
	return &aTMessageRepo{data: data}
}

func (atr *aTMessageRepo) CareerData(data *dto.WSATMessageData) (*dto.MessageToCreate, error) {
	return &dto.MessageToCreate{
		Content: "数据民测试",
		MsgID:   data.ID,
		MessageReference: &dto.MessageReference{
			MessageID:             data.ID,
			IgnoreGetMessageError: false,
		},
	}, nil
}
