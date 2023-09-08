package main

import (
	"relation/conf"
	"relation/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	r := gin.Default()
	re := r.Group("/douyin/relation/")
	re.GET("follow/list", controllers.Follow)            //关注列表
	re.GET("follower/list", controllers.Follower)        //粉丝列表
	re.GET("friend/list", controllers.Friend)            //好友列表
	re.GET("message/list", controllers.Message)          //消息列表
	re.POST("action/follow", controllers.ActionFollow)   //关注操作
	re.POST("action/message", controllers.ActionMessage) //发送消息
	r.Run(":8080")
}
