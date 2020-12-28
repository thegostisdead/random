package main

import (
	"fmt"
	"regexp"
)

type chapter struct {
	title      string
	folderName string
	url        string
}

func extractPicturesFromHTML(html string) []string {
	var res []string
	var re = regexp.MustCompile(`(?sm)\<img\sclass\=\"aligncenter\ssize\-full\swp\-image\-.*?\"\ssrc\=\"(https://scantoon\.com\/wp\-content\/uploads\/.*?.jpg)[\"]`)
	for _, match := range re.FindAllStringSubmatch(html, -1) {
		res = append(res, match[1])
	}

	return res
}

func getPictureFilename(url string) string {
	var re = regexp.MustCompile(`(?ms)/[0-9][0-9]\/(.*?\.jpg)`)

	for _, match := range re.FindAllStringSubmatch(url, -1) {
		return match[1]
	}
	return "error"

}

func downloadChapter(chapter chapter) {

	fmt.Printf("Downloading chapter %s\n", chapter.title)

	chapterHTML := downloadHTML(chapter.url)

	var imgLinks = extractPicturesFromHTML(chapterHTML)

	//fmt.Printf("%s\n", imgLinks)
	var webtoonFolder = "output/" + getNameWebToon(chapter.title)
	createFolder(webtoonFolder + "/" + chapter.folderName)

	for _, image := range imgLinks {
		path := webtoonFolder + "/" + chapter.folderName+"/"+getPictureFilename(image)
		_ = DownloadFile(image, path)
	}

}
