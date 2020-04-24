package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

var roamingDir string = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH") + "\\AppData\\Roaming\\test"

type ModpackConfig struct {
	Forge  string   `json:"forge"`
	Mods   []string `json:"mods"`
	Server string   `json:"server"`
}

func getConfig() *ModpackConfig {
	config := &ModpackConfig{}
	data, err := ioutil.ReadFile("modpack.json")
	checkError(err)
	err = json.Unmarshal(data, config)
	checkError(err)
	return config
}

type Installer struct {
	downloader *Downloader
	modpack    *ModpackConfig
	dir        string
}

func newInstaller() *Installer {
	return &Installer{
		downloader: newDownloader("./server"),
		modpack:    getConfig(),
		dir:        "./server",
	}
}
func (installer Installer) install() {
	mods := installer.modpack.Mods
	forge := installer.modpack.Forge
	downloads := append(mods, forge)

	installer.downloader.downloadInMemory(downloads...)

	fmt.Print("Installing mods....")
	go installer.downloader.saveFromMemory(installer.dir)

	if *server {

		installer.downloader.downloadFiles(installer.modpack.Server)
		installer.generateEula()
		fmt.Print("Executing server...")
		installer.executeServer()
	}

}

func (installer Installer) generateEula() {
	file, err := os.Create(installer.dir + "/" + "eula.txt")
	checkError(err)
	data := []byte("eula=true")
	file.Write(data)
}
func (installer Installer) executeServer() {

	cmd := exec.Command("java", "-Xmx2G", "-Xms1G", "-jar", "./server/server.jar", "nogui")
	io.Pipe()

	output, err := cmd.CombinedOutput()
	checkError(err)
	go fmt.Printf("%s\n", output)

	err = cmd.Run()

	checkError(err)

}
