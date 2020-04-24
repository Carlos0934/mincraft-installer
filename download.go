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

type Downloader struct {
	client *http.Client

	files   map[string][]byte
	dir     string
	printer *DownloadPrinter
	wg      *sync.WaitGroup
	counter uint8
}

func newDownloader(dir string) *Downloader {
	client := http.Client{}

	return &Downloader{
		client: &client,
		dir:    dir,
		files:  make(map[string][]byte),

		printer: &DownloadPrinter{},
		wg:      &sync.WaitGroup{},
	}
}

func (downloader *Downloader) downloadFiles(urls ...string) {

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

func (downloader *Downloader) downloadInMemory(urls ...string) {
	downloader.wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			downloader.counter++
			defer downloader.wg.Done()

			req, err := downloader.client.Get(url)
			checkError(err)
			defer req.Body.Close()
			reader := io.TeeReader(req.Body, downloader.printer)

			filename := path.Base(req.Request.URL.Path)
			downloader.saveInMemory(filename, reader)

			downloader.counter--
		}(url)
	}

	downloader.wg.Wait()

	fmt.Println("\nAll downloads finished succesfully")
}

func (downloader *Downloader) saveInMemory(filename string, writter io.Reader) {
	data, err := ioutil.ReadAll(writter)
	checkError(err)
	downloader.files[filename] = data
}

func (downloader *Downloader) saveFromMemory(dirs ...string) {
	downloader.wg.Add(len(dirs) * len(downloader.files))
	for _, dir := range dirs {

		for filename, data := range downloader.files {
			go downloader.saveBytes(dir, filename, data)

		}
	}
	downloader.wg.Wait()
	fmt.Println("All downloads have been installed")

}
func (downloader *Downloader) saveBytes(dir string, filename string, data []byte) {

	file, err := os.Create(dir + "/" + filename)
	checkError(err)
	_, err = file.Write(data)
	checkError(err)

	file.Close()
	downloader.wg.Done()
}

func (downloader *Downloader) clear() {
	downloader.files = make(map[string][]byte)
}

func (downloader *Downloader) save(filename string, writter io.Reader) {
	dir := downloader.dir + "/" + filename
	file, err := os.Create(dir)

	checkError(err)
	defer file.Close()

	_, err = io.Copy(file, writter)

	checkError(err)

}
