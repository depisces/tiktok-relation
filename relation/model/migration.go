package model

func migration() {
	DB.AutoMigrate(&User{}, &Follow{}, &Friend{}, &Message{})
}
