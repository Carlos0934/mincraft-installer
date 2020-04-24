package main

import "fmt"

func main() {
	modsLoader := newDownloader(10)
	modsLoader.download("https://media.forgecdn.net/files/2894/944/RebornCore-1.12.2-3.19.1.521-universal.jar")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
