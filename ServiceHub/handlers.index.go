package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func showIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"title": "Home",
		},
	)
}
