package main

import (
	"encoding/json"
	"io/ioutil"
)

type ModpackConfig struct {
	Forge string   `json:"forge"`
	Mods  []string `json:"mods"`
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
}

func (installer Installer) install() {

}
