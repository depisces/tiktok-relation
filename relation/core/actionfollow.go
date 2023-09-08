package core

import (
	"relation/model"
	"relation/service"
)

// actionType-- 1-关注  2-取关
func BuildProtoActionfollow(followingUserId int64, followedUserId int64, actionType int64) *service.DouyinRelationActionFollowResponse {
	code := int32(0)
	msg := "ok"
	if actionType == 1 { //关注操作
		//数据库操作
		err := model.CreateFollowByUserId(followingUserId, followedUserId)
		if err != nil {
			code = 1
			msg = err.Error()
		}
	} else { //取关操作
		//数据库操作
		err := model.DeleteFollowByUserId(followingUserId, followedUserId)
		if err != nil {
			code = 1
			msg = err.Error()
		}
	}
	ret := service.DouyinRelationActionFollowResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
	return &ret
}
