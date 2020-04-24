package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	installer := newInstaller()

	installer.install()
	for true {

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
