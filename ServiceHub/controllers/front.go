package controllers

import (
	"gin_session/helpers"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	userName := helpers.GetUserName(c)

	c.HTML(200, "login.html", gin.H{
		"userName":    userName,
		"currentPage": "login",
		"success":     helpers.GetFlashCookie(c, "success"),
	})
}
