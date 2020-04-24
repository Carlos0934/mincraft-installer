package main

import "fmt"

func main() {
	config := getConfig()
	fmt.Println(config.Mods)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
