package main

import (
	"github.com/gin-gonic/gin"
)


func AuthRequiredMiddleware(c *gin.Context) {
	/*
	userID, userIDErr := primitive.ObjectIDFromHex(c.GetHeader("X-User-Id"))
	authToken := c.GetHeader("X-Auth-Token")

	if userIDErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessageResponse("You must be logged in to do this."))
		return
	}

	_, authErr := GetUserByToken(userID, authToken)

	if authErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessageResponse("You must be logged in to do this."))
		return
	}

	// set userID
	c.Set("userID", userID)

	// Pass on to the next-in-chain
	c.Next()
	*/

}