package gozip

import (
	"github.com/nareix/curl"
	"log"
	"strings"
	"time"
)

func DownloadFile(url string, fileName string) {
	req := curl.New(url)

	req.Method("POST")
	req.SaveToFile(fileName)

	// Print progress status per one second
	req.Progress(func(p curl.ProgressStatus) {
		log.Println(
			"speed", curl.PrettySpeedString(p.Speed),
			"len", curl.PrettySizeString(p.ContentLength),
			"got", curl.PrettySizeString(p.Size),
		)
	}, time.Second)

	req.Do()
}

func DownloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	req := curl.New(url)

	req.Method("POST")
	req.SaveToFile(fileName)

	// Print progress status per one second
	req.Progress(func(p curl.ProgressStatus) {
		log.Println(
			"speed", curl.PrettySpeedString(p.Speed),
			"len", curl.PrettySizeString(p.ContentLength),
			"got", curl.PrettySizeString(p.Size),
		)
	}, time.Second)

	req.Do()
}
