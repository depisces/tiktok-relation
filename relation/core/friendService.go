package core

import (
	"fmt"
	"relation/model"
	"relation/service"
)

func BuildProtoFriend(id int64, myUserId int64) *service.DouyinRelationFriendListResponse {
	//返回数据库实体
	friend, err := model.FindFriendByUserId(id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	var userlist []*service.FriendUser
	for _, v := range *friend {
		u, _ := model.FindUserById(v.FriendUserId)
		userlist = append(userlist, &service.FriendUser{User: BuildProtoUser(u, myUserId), Message: "", MsgType: 1})
	}
	friendResp := service.DouyinRelationFriendListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   userlist,
	}
	return &friendResp
}
