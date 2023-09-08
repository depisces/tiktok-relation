package controllers

import (
	"relation/core"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Message(c *gin.Context) {
	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))
	token, _ := strconv.Atoi(c.Query("token"))
	myUserId := core.GetIdByToken(int64(token))
	from_user_id := myUserId
	list := core.BuildProtoMessage(int64(to_user_id), int64(from_user_id))
	c.JSON(200, list)
}
