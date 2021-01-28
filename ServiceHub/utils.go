package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Users struct {
	Host     string `form:"id"`
	UserName string `form:"username"`
	Password string `form:"password"`
	Token    string ""
}

type Token struct {
	Jwt string
}

const authHeader string = "Bearer "

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

func get(url string, user *Users) string {
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, payload)
	fmt.Println(authHeader + user.Token)
	req.Header.Add("Authorization", authHeader+user.Token)

	response, err := client.Do(req)
	fmt.Println("Response : ")
	fmt.Println(response)
	if err != nil {
		fmt.Printf("Error %s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return string(contents)
	}
	return "Error"
}
