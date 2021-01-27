package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.Static("/assets", "./static")

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	router.GET("/login", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,

			"login.html",

			gin.H{
				"title": "Login",
			},
		)

	})

	router.GET("/home", func(c *gin.Context) {

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
