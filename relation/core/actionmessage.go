package core

import (
	"relation/model"
	"relation/service"
)

// actionType--  1-发送消息
func BuildProtoActionMessage(fromUserId int64, toUserId int64, actionType int64, content string) *service.DouyinRelationActionMessageResponse {
	code := int32(0)
	msg := "ok"
	if actionType == 1 {
		err := model.CreateMessageByUserIdAndContent(fromUserId, toUserId, content)
		if err != nil {
			code = 1
			msg = err.Error()
		}
	}
	ret := service.DouyinRelationActionMessageResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
	return &ret
}
