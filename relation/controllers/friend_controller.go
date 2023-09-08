package controllers

import (
	"relation/core"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Friend(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	token, _ := strconv.Atoi(c.Query("token"))
	myUserId := core.GetIdByToken(int64(token))

	list := core.BuildProtoFriend(int64(id), myUserId)
	c.JSON(200, list)
}
