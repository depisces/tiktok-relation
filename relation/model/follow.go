package model

import (
	"errors"
	"sync"
)

type Follow struct {
	Id int64 `gorm:"primary_key"`
	//User            User `gorm:"foreignKey:FollowingUserId FollowedUserId;references:Id"`
	FollowingUserId int64 `gorm:""`
	FollowedUserId  int64 `gorm:""`
}

func (Follow) TableName() string {
	return "follow"
}

type FollowDao struct {
}

var followDao *FollowDao
var followOnce sync.Once

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

// 用关注用户的id查询
func FindFollowByFollowingUserId(id int64) (*[]Follow, error) {
	var ret *[]Follow
	d := DB.Where("following_user_id=?", id).Find(&ret)
	err := d.Error
	if err != nil {
		return nil, err
	}
	return ret, err
}

// 用被关注人的id查询
func FindFollowByFollowedUserId(id int64) (*[]Follow, error) {
	var ret *[]Follow
	d := DB.Where("followed_user_id=?", id).Find(&ret)
	err := d.Error
	if err != nil {
		return nil, err
	}
	return ret, err
}

// 关注操作
func CreateFollowByUserId(followingUserId int64, followedUserId int64) error {
	if CheckIsFollowedByUserId(followingUserId, followedUserId) {
		return errors.New("您已关注该用户")
	}

	//添加记录
	k := Follow{FollowingUserId: followingUserId, FollowedUserId: followedUserId}
	d := DB.Create(&k)

	if CheckIsFollowedByUserId(followedUserId, followingUserId) {
		//如果对方关注了自己，则添加好友关系
		CreateFriendByUserId(followingUserId, followedUserId)
	}

	err := d.Error
	if err != nil {
		return err
	}
	return nil
}

// 取关操作
func DeleteFollowByUserId(followingUserId int64, followedUserId int64) error {
	if !CheckIsFollowedByUserId(followingUserId, followedUserId) {
		//未关注该用户
		return errors.New("你未关注该用户")
	}
	if CheckIsFollowedByUserId(followedUserId, followingUserId) {
		//对方关注了自己，则取消好友关系
		DeleteFriendByUserId(followingUserId, followedUserId)
	}
	result := DB.Where("following_user_id=? and followed_user_id=?", followingUserId, followedUserId).Delete(&Follow{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 检查fromUser是否关注了toUser
func CheckIsFollowedByUserId(fromUserId int64, toUserId int64) bool {
	result := DB.Where("following_user_id=? and followed_user_id=?", fromUserId, toUserId).First(&Follow{})
	return result.RowsAffected != 0
}
