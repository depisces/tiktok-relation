package model

import (
	"sync"
	"time"
)

type Message struct {
	Id         int64 `gorm:"primary_key"`
	ToUserId   int64
	FromUserId int64
	Content    string
	CreateTime time.Time
}

func (Message) TableName() string {
	return "message"
}

type MessageDao struct{}

var messageDao *MessageDao
var messageOnce sync.Once

func NewMessageDaoInstance() *MessageDao {
	messageOnce.Do(func() {
		messageDao = &MessageDao{}
	})
	return messageDao
}

// 用对话双方的id查找消息列表
func FindMessageByUserId(id1 int64, id2 int64) (*[]Message, error) {
	var message *[]Message
	d := DB.Where("to_user_id=? and from_user_id=?", id1, id2).Or("to_user_id=? and from_user_id=?", id2, id1).Order("create_time asc").Find(&message)
	err := d.Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

// 生成消息
func CreateMessageByUserIdAndContent(fromUserId int64, toUserId int64, con string) error {
	result := DB.Create(&Message{FromUserId: fromUserId, ToUserId: toUserId, Content: con, CreateTime: time.Now()})
	return result.Error
}
