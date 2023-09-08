package core

import (
	"relation/model"
	"relation/service"
)

func BuildProtoMessage(id1 int64, id2 int64) *service.DouyinMessageChatResponse {
	//返回数据库实体
	message, err := model.FindMessageByUserId(id1, id2)
	if err != nil {
		return nil
	}
	//消息列表
	var messagelist []*service.Message
	for _, v := range *message {
		m := service.Message{
			Id:         v.Id,
			ToUserId:   v.ToUserId,
			FromUserId: v.FromUserId,
			Content:    v.Content,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		messagelist = append(messagelist, &m)
	}

	messageResp := service.DouyinMessageChatResponse{
		StatusCode:  0,
		StatusMsg:   "ok",
		MessageList: messagelist,
	}
	return &messageResp
}
