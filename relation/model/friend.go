package model

import "sync"

type Friend struct {
	Id           int64 `gorm:"primary_key"`
	UserId       int64
	FriendUserId int64
}

func (Friend) TableName() string {
	return "Friend"
}

type FriendDao struct{}

var friendDao *FriendDao
var friendOnce sync.Once

func NewFriendDaoInstance() *FriendDao {
	friendOnce.Do(
		func() {
			friendDao = &FriendDao{}
		})
	return friendDao
}

//查找该用户的好友
func FindFriendByUserId(id int64) (*[]Friend, error) {
	var ret *[]Friend
	d := DB.Where("friend_user_id").Find(&ret)
	err := d.Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// 创建好友关系
func CreateFriendByUserId(userId int64, friendUserId int64) error {
	//好友关系是双向的
	DB.Create(&Friend{UserId: userId, FriendUserId: friendUserId})
	DB.Create(&Friend{UserId: friendUserId, FriendUserId: userId})
	return nil
}

// 删除好友关系
func DeleteFriendByUserId(userId int64, friendUserId int64) error {
	DB.Where("user_id=? and friend_user_id=?", userId, friendUserId).Delete(&Friend{})
	DB.Where("user_id=? and friend_user_id=?", friendUserId, userId).Delete(&Friend{})
	return nil
}
