package main

import (
	"net/http"

	"gin_session/controllers"
	"gin_session/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func Private() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Request.Cookie(helpers.CookieSessionName)

		if err != nil {
			c.Redirect(302, "/login")
			c.Abort()
		}
	}
}

func main() {
	// Initialisation du routeur
	r := gin.Default()

	r.Static("/assets", "./static")
	r.LoadHTMLGlob("views/*")

	r.GET("/", controllers.IndexHandler)

	admin := r.Group("/home")
	admin.Use(Private())
	{

		admin.GET("", controllers.HomeHandler)

	}

	r.GET("/login", controllers.LoginHandlerForm)
	r.POST("/login", controllers.LoginHandler)

	r.GET("/logout", controllers.LogoutHandler)

	// Port du serveur
	//r.Run(":3000")
	http.ListenAndServe(":8080",
		csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(r))
}
