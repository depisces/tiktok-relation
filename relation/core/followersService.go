package core

import (
	"relation/model"
	"relation/service"
)

func BuildProtoFollowers(id int64, myUserId int64) *service.DouyinRelationFollowerListResponse {
	//返回数据库实体
	follow, err := model.FindFollowByFollowedUserId(id)
	if err != nil {
		return nil
	}

	//生成用户列表
	var userlist []*service.User
	for _, v := range *follow { //遍历该用户粉丝的id
		u, _ := model.FindUserById(v.FollowingUserId)
		userlist = append(userlist, BuildProtoUser(u, myUserId))
	}

	followResp := service.DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   userlist,
	}
	return &followResp
}
