package gozip

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/nareix/curl"
	"github.com/vaughan0/go-ini"
	"log"
	"os"
	"strings"
	"time"
)

func DownloadFile(url string, fileName string, verbose bool) {
	req := curl.New(url)

	req.Method("POST")
	req.SaveToFile(fileName)

	// Print progress status per one second
	req.Progress(func(p curl.ProgressStatus) {
		if verbose {
			log.SetOutput(os.Stdout)
			log.Println(
				"speed", curl.PrettySpeedString(p.Speed),
				"len", curl.PrettySizeString(p.ContentLength),
				"got", curl.PrettySizeString(p.Size),
			)
		}
	}, time.Second)

	req.Do()
}

func DownloadFromUrl(url string, verbose bool) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	req := curl.New(url)

	req.Method("POST")
	req.SaveToFile(fileName)

	// Print progress status per one second
	req.Progress(func(p curl.ProgressStatus) {
		if verbose {
			log.SetOutput(os.Stdout)
			log.Println(
				"speed", curl.PrettySpeedString(p.Speed),
				"len", curl.PrettySizeString(p.ContentLength),
				"got", curl.PrettySizeString(p.Size),
			)
		}
	}, time.Second)

	req.Do()
}

func GetIni(inifile, section, name string) string {
	cfg, err := ini.LoadFile(inifile)
	if err != nil {
		log.Fatal(err)
	}
	getname, ok := cfg.Get(section, name)
	if !ok {
		log.Fatal("app not found in file INI ", inifile)
	}
	return getname
}

func GoQueryGet(url, find1, find2 string) string {
	var fileName string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(find1).Each(func(i int, s *goquery.Selection) {
		fileName = s.Find(find2).Text()
	})

	return fileName
}

func LogToFile(logfile, msg string, verbose bool) {
	if verbose {
		log.SetOutput(os.Stdout)
		log.Println(msg)
	}
	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(msg)
}
