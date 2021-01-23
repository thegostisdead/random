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

func get(url string, portainer *Portainer) string {
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, payload)
	fmt.Println(authHeader + portainer.Token)
	req.Header.Add("Authorization", authHeader+portainer.Token)

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
