package controllers

import (
	"relation/core"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActionFollow(c *gin.Context) {
	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	token, _ := strconv.Atoi(c.Query("token"))
	myUserId := core.GetIdByToken(int64(token))
	list := core.BuildProtoActionfollow(int64(myUserId), int64(to_user_id), int64(action_type))
	c.JSON(200, list)
}
