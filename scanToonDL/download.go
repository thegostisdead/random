package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func processMatchedURL(matched string) string {

	startRe := regexp.MustCompile(`(?msU)(<option.*?value=".*?".*?data-url=").*?`)
	endRe := regexp.MustCompile(`(?ms).*?(\".*?>.*<\/option>)`)

	start := startRe.ReplaceAllString(matched, "$1")
	startRemoved := strings.ReplaceAll(matched, start, "")

	end := endRe.ReplaceAllString(startRemoved, "$1")
	endRemoved := strings.ReplaceAll(startRemoved, end, "")

	return endRemoved

}

func extractChaptersFromHTML(html string) []string {
	var res []string

	var re = regexp.MustCompile(`(?m)\<option
							
							value\=\".*?\"
							data\-url\=\"(.*?)\"
														\>.*?						\<\/option\>
`)

	for _, match := range re.FindAllString(html, -1) {
		res = append(res, processMatchedURL(match))
	}
	fmt.Println(res)
	return res

}

func getTitleFromUrl(url string) string {

	res := strings.ReplaceAll(url, "https://scantoon.com/scans/", "")
	res = strings.ReplaceAll(res, "-", " ")

	return strings.Trim(res, "/")

}

func getFolderNameFromUrl(url string) string {
	re := regexp.MustCompile(`(?ms)\d+`)
	for _, match := range re.FindAllString(url, -1) {
		return match
	}

	return "Error"
}

func getAllChapters(url string) []chapter {
	var chapters []chapter
	var htmlSource = downloadHTML(url)
	extractedLinkChapters := extractChaptersFromHTML(htmlSource)

	for _, chap := range extractedLinkChapters {
		title := getTitleFromUrl(chap)
		folderName := getFolderNameFromUrl(chap)

		chapters = append(chapters, chapter{
			title:      title,
			folderName: folderName,
			url:        chap,
		})
	}

	return chapters
}

func getNameWebToon(title string) string {
	var re = regexp.MustCompile(`(?m)(.*?)scan\s.*?vf`)
	for _, match := range re.FindAllStringSubmatch(title, -1) {
		return strings.TrimRight(match[1], " ")
	}
	return "Error"
}

func download(webtoonURL string) {

	if _, err := os.Stat("output"); os.IsNotExist(err) {
		fmt.Println("Create output directory...")
		createFolder("output")
	}

	inputChapter := chapter{
		title:      getTitleFromUrl(webtoonURL),
		folderName: getFolderNameFromUrl(webtoonURL),
		url:        webtoonURL}

	chaptersURL := getAllChapters(webtoonURL)

	chaptersURL = append(chaptersURL, inputChapter)

	createFolder("output/" + getNameWebToon(chaptersURL[0].title)) // create folder to store all chapters

	for _, chapter := range chaptersURL {
		downloadChapter(chapter)
	}

}
