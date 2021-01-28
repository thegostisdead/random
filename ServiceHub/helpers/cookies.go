package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"net/http"
	"net/url"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

var CookieSessionName = "session"

func SetSession(userName string, c *gin.Context) {
	value := map[string]string{
		"name": userName,
	}

	if encoded, err := cookieHandler.Encode(CookieSessionName, value); err == nil {
		cookie := &http.Cookie{
			Name:     CookieSessionName,
			Value:    encoded,
			Path:     "/",
			HttpOnly: true,
			//Secure: true, Only if  HTTPS
		}
		http.SetCookie(c.Writer, cookie)
	}
}

func GetUserName(c *gin.Context) (userName string) {
	if cookie, err := c.Request.Cookie(CookieSessionName); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(CookieSessionName, cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}

	return userName
}

func ClearSession(c *gin.Context) {
	cookie := &http.Cookie{
		Name:   CookieSessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(c.Writer, cookie)
}

func encode(value string) string {
	encode := &url.URL{Path: value}
	return encode.String()
}

func decode(value string) string {
	decode, _ := url.QueryUnescape(value)
	return decode
}

func SetFlashCookie(c *gin.Context, name string, value string) {
	cookie := &http.Cookie{
		Name:   name,
		Value:  encode(value),
		Path:   "/",
		MaxAge: 1,
	}

	http.SetCookie(c.Writer, cookie)
}

func GetFlashCookie(c *gin.Context, name string) (value string) {
	cookie, err := c.Request.Cookie(name)

	var cookieValue string
	if err == nil {
		cookieValue = cookie.Value
	} else {
		cookieValue = cookieValue
	}

	return decode(cookieValue)
}
