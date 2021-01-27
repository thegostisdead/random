package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login",
		},
	)
}