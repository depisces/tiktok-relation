package core

import (
	"relation/model"
	"relation/service"
)

type UserService struct {
}

func BuildProtoUser(item *model.User, myUserId int64) *service.User {
	user := service.User{
		Id:              item.Id,
		Name:            item.Name,
		FollowCount:     item.FollowingCount,
		FollowerCount:   item.FollowersCount,
		IsFollow:        model.CheckIsFollowedByUserId(myUserId, item.Id),
		Avatar:          item.Avatar,
		BackgroundImage: item.BackgroundImage,
		Signature:       item.Signature,
		TotalFavorited:  item.TotalFavorited,
		WorkCount:       item.WorkCount,
		FavoriteCount:   item.FavoriteCount,
	}
	return &user
}
