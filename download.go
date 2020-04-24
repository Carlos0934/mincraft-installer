package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
)

const (
	OutMemory int = 0
	InMemory  int = 1
)

type Downloader struct {
	client *http.Client

	files   map[string][]byte
	dir     string
	printer *DownloadPrinter
	wg      *sync.WaitGroup
	counter uint8
}

func newDownloader(max uint8) *Downloader {
	client := http.Client{}

	return &Downloader{
		client: &client,
		dir:    "./tmp/",
		files:  make(map[string][]byte),

		printer: &DownloadPrinter{},
		wg:      &sync.WaitGroup{},
	}
}

func (downloader *Downloader) downloadFiles(download int, urls ...string) {

	downloader.wg.Add(len(urls))
	for _, url := range urls {

		go downloader.download(url)
	}

	downloader.wg.Wait()
	fmt.Println("\nAll downloads finished succesfully")
}

func (downloader *Downloader) download(url string) {
	downloader.counter++
	defer downloader.wg.Done()

	req, err := downloader.client.Get(url)
	checkError(err)
	defer req.Body.Close()
	reader := io.TeeReader(req.Body, downloader.printer)

	filename := path.Base(req.Request.URL.Path)
	downloader.save(filename, reader)

	downloader.counter--
}

func (downloader *Downloader) saveInMemory(filename string, writter io.Reader) {
	data, err := ioutil.ReadAll(writter)
	checkError(err)
	downloader.files[filename] = data
}
func (downloader *Downloader) saveFromMemory(dirs ...string) {
	for _, dir := range dirs {

		for filename, data := range downloader.files {
			go downloader.saveBytes(dir, filename, data)

		}
	}

}
func (Downloader) saveBytes(dir string, filename string, data []byte) {
	file, err := os.Create(dir + "/" + filename)
	checkError(err)
	_, err = file.Write(data)
	checkError(err)

	file.Close()
}

func (downloader *Downloader) clear() {
	downloader.files = make(map[string][]byte)
}

func (downloader *Downloader) save(filename string, writter io.Reader) {
	dir := downloader.dir + filename
	file, err := os.Create(dir)

	checkError(err)
	defer file.Close()

	_, err = io.Copy(file, writter)

	checkError(err)

}
