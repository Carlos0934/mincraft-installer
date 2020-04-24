package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

type Downloader struct {
	client  *http.Client
	actives chan (uint8)
	max     uint8
	files   map[string][]byte
	dir     string
}

func newDownloader(max uint8) *Downloader {
	client := http.Client{}

	return &Downloader{
		client:  &client,
		dir:     "./",
		files:   make(map[string][]byte),
		actives: make(chan uint8),
		max:     max,
	}
}

func (downloader *Downloader) downloadFiles(urls ...string) {
	for _, url := range urls {
		go downloader.download(url)
	}

}

func (downloader *Downloader) download(url string) {

	first := time.Now()

	req, err := downloader.client.Get(url)
	checkError(err)
	defer req.Body.Close()
	file, err := ioutil.ReadAll(req.Body)

	checkError(err)

	fmt.Println(time.Now().Sub(first).Seconds())
	filename := path.Base(req.Request.URL.Path)
	downloader.save(filename, file)
}

func (downloader *Downloader) save(filename string, data []byte) {

	err := ioutil.WriteFile(downloader.dir+filename, data, os.ModeAppend)

	checkError(err)

}
