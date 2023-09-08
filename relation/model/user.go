package model

import "sync"

type User struct {
	Id              int64  `gorm:"primary_key"`
	Name            string `gorm:"unique"`
	FollowingCount  int64  `gorm:"default:0"`
	FollowersCount  int64  `gorm:"default:0"`
	Password        string `gorm:"default:'-'"`
	Avatar          string `gorm:"default:'-'"`
	BackgroundImage string `gorm:"default:'-'"`
	Signature       string `gorm:"default:'-'"`
	TotalFavorited  int64  `gorm:"default:0"`
	WorkCount       int64  `gorm:"default:0"`
	FavoriteCount   int64  `gorm:"default:0"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct{}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func FindUserById(id int64) (*User, error) {
	user := User{Id: id}
	result := DB.Where("id=?", id).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
func FindUserByName(name string) (*User, error) {
	user := User{Name: name}
	result := DB.Where("name=?", name).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
