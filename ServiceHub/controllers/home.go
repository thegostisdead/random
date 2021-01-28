package controllers

import (
	"gin_session/helpers"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	var users []Users

	userName := helpers.GetUserName(c)

	c.HTML(200, "home.html", gin.H{
		"userName":    userName,
		"currentPage": "home",
		"users":       users,
		"success":     helpers.GetFlashCookie(c, "success"),
	})
}
