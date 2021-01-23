package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Portainer struct {
	Host     string `json:"host"`
	User     string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Data     string `json:"data"`
	Settings string `json:"settings"`
}

type Token struct {
	Jwt string
}

const apiPath string = "/api"
const apiAuthPath string = "/auth"
const apiEndpoints string = "/endpoints/1/docker/containers/json?all=1"
const authHeader string = "Bearer "
const portainerSettings string = "/settings"

var portainer Portainer

func getPortainerHandler(w http.ResponseWriter, r *http.Request) {
	portainerListBytes, err := json.Marshal(portainer)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(portainerListBytes)
}

func getToken(portainer *Portainer) string {
	toRequestURL := "http://" + portainer.Host + apiPath + apiAuthPath
	fmt.Println(toRequestURL)

	postBody, _ := json.Marshal(map[string]string{
		"Username": portainer.User,
		"Password": portainer.Password,
	})
	responseBody := bytes.NewBuffer(postBody)

	resBytes := post(toRequestURL, "application/json", responseBody)

	var token Token

	_ = json.Unmarshal(resBytes, &token)

	fmt.Printf("TOKEN = %v\n", token.Jwt)
	return token.Jwt
}

func updateToken(portainer *Portainer) {
	portainer.Token = getToken(portainer)
	fmt.Printf("Updating token ....")

}

func fetchPortainerSettings(portainer *Portainer) {
	fmt.Println("Fetching portianer settings ....")
	resp := get("http://"+portainer.Host+apiPath+portainerSettings, portainer)
	fmt.Println(resp)
	portainer.Settings = resp
}

func fetchContainerData(portainer *Portainer) {
	fmt.Println("Fetching portianer data ....")
	fmt.Println("http://" + portainer.Host + apiPath + apiEndpoints)
	resp := get("http://"+portainer.Host+apiPath+apiEndpoints, portainer)
	fmt.Println(resp)
	portainer.Data = resp
}

func createPortainerHandler(w http.ResponseWriter, r *http.Request) {
	resPortainer := Portainer{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resPortainer.Host = r.Form.Get("host")
	resPortainer.User = r.Form.Get("username")
	resPortainer.Password = r.Form.Get("password")
	resPortainer.Token = getToken(&resPortainer)
	fmt.Println("Ok jusque ala ")
	fetchContainerData(&resPortainer)
	fetchPortainerSettings(&resPortainer)

	portainer = resPortainer

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
