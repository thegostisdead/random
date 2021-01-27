package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine



func main() {


	router := gin.Default()

	router.Static("/assets", "./static")

	router.LoadHTMLGlob("templates/*")

	api := router.Group("/api")
	{
		api.GET("/books", AuthRequiredMiddleware, func(c *gin.Context) {
			userID := c.MustGet("userID")

			c.JSON(200, gin.H{
				"userId": userID,
			})
		})
	}

	initializeRoutes()


	router.GET("/home", AuthRequiredMiddleware,func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"home.html",
			gin.H{
				"title": "Home",
			},
		)

	})

	// Start serving the application
	router.Run()

}
