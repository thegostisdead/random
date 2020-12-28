package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func welcome() {
	fmt.Println("ScanToonDL, a simple downloader...")
	fmt.Println("Paste your webtoon URL and this script will download all the chapters for you.")
}

func handleURL() *string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter url to your webtoon: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Your enter : %s\n", text)

	if text != "" {
		var res = strings.Trim(text, "\n")
		return &res
	} else {
		println("Error, you type an empty url!")
		return nil
	}
}

func main() {
	welcome()

	var urlInput = handleURL()


	if isURLValid(*urlInput) {
		download(*urlInput)
		fmt.Println("Download finished")
		os.Exit(0)
	} else {
		panic("Error your webtoon url is not valid, please try again")
	}

}
