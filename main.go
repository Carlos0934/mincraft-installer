package main

import "fmt"

func main() {
	modsLoader := newDownloader(1)
	modsLoader.downloadFiles("https://media.forgecdn.net/files/2894/944/RebornCore-1.12.2-3.19.1.521-universal.jar", "https://media.forgecdn.net/files/2518/667/Baubles-1.12-1.5.2.jar",
		"https://files.minecraftforge.net/maven/net/minecraftforge/forge/1.12.2-14.23.5.2847/forge-1.12.2-14.23.5.2847-installer.jar",
		"https://media.forgecdn.net/files/2920/254/astralsorcery-1.12.2-1.10.23.jar")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
