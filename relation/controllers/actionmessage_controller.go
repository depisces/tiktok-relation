package controllers

import (
	"relation/core"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActionMessage(c *gin.Context) {
	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	token, _ := strconv.Atoi(c.Query("token"))
	myUserId := core.GetIdByToken(int64(token))

	//POST表单接收聊天内容
	//content := c.PostForm("content")

	//参数接收聊天内容
	content := c.Query("content")

	list := core.BuildProtoActionMessage(myUserId, int64(to_user_id), int64(action_type), content)

	c.JSON(200, list)
}
