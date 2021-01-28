package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin_session/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Users struct {
	Host     string `form:"host"`
	UserName string `form:"username"`
	Password string `form:"password"`
	Token    string ""
}

type Token struct {
	Jwt string
}

const apiPath string = "/api"
const apiAuthPath string = "/auth"
const apiEndpoints string = "/endpoints/1/docker/containers/json?all=1"
const portainerSettings string = "/settings"

func LoginHandlerForm(c *gin.Context) {
	// Récupération du nom d'utilisateur pour le templating
	userName := helpers.GetUserName(c)

	c.HTML(200, "login.html", gin.H{
		"userName":       userName,
		"currentPage":    "login",
		"warning":        helpers.GetFlashCookie(c, "warning"),
		csrf.TemplateTag: csrf.TemplateField(c.Request),
	})
}

func post(endpoint string, contentType string, responseBody *bytes.Buffer) []byte {

	resp, err := http.Post(endpoint, contentType, responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		log.Fatalln(err)
	}
	return body
}

func LoginHandler(c *gin.Context) {
	var user Users

	if c.Bind(&user) == nil {

		clearPassword := user.Password
		hostname := user.Host
		username := user.UserName

		fmt.Println(clearPassword)
		fmt.Println(hostname)
		fmt.Println(username)
		token := getToken(&user)

		if token != "" {
			if strings.Contains(token, "Unauthorized") {
				helpers.SetFlashCookie(c, "warning", "Mot de passe incorrect")
				c.Redirect(302, "/login")
			} else {
				user.Token = token
				helpers.SetSession(user.UserName, c)
				c.Redirect(302, "/home")
			}
		} else {
			helpers.SetFlashCookie(c, "warning", "Mot de passe incorrect")
			c.Redirect(302, "/login")
		}

	} else {
		// Champs non remplis
		helpers.SetFlashCookie(c, "warning", "Champs non remplis")
		c.Redirect(302, "/login")
	}
}

func LogoutHandler(c *gin.Context) {
	helpers.ClearSession(c)
	helpers.SetFlashCookie(c, "success", "Vous êtes désormais déconnecté(e)")
	c.Redirect(302, "/")
}

func getToken(user *Users) string {
	toRequestURL := "http://" + user.Host + apiPath + apiAuthPath
	fmt.Println(toRequestURL)

	postBody, _ := json.Marshal(map[string]string{
		"Username": user.UserName,
		"Password": user.Password,
	})
	responseBody := bytes.NewBuffer(postBody)

	resBytes := post(toRequestURL, "application/json", responseBody)

	var token Token

	_ = json.Unmarshal(resBytes, &token)

	fmt.Printf("TOKEN = %v\n", token.Jwt)
	return token.Jwt
}
