package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func downloadHTML(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(html)
}


func isURLValid(url string) bool {
	if strings.HasPrefix(url, "https://scantoon.com/scans/") {
		return true
	}
	return false
}



func createFolder(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}