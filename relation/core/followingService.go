package core

import (
	"relation/model"
	"relation/service"
)

func BuildProtoFollowing(id int64, myUserId int64) *service.DouyinRelationFollowListResponse {
	//返回数据库实体
	follow, err := model.FindFollowByFollowingUserId(id)
	if err != nil {
		return nil
	}

	//生成对应用户列表
	var userlist []*service.User
	for _, v := range *follow {
		u, _ := model.FindUserById(v.FollowedUserId)
		userlist = append(userlist, BuildProtoUser(u, myUserId))
	}
	followResp := service.DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   userlist,
	}
	return &followResp
}
