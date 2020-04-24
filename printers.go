package main

import (
	"fmt"
	"strings"
)

type DownloadPrinter struct {
}

func (DonwloadPrinter) Write(file []byte) (int, error) {
	fmt.Printf("\r%s", strings.Repeat(" ", 50))

	fmt.Printf("\rDownloading... %s complete")
}

func (DonwloadPrinter) mbSize(file []byte) float32 {
	binary.size()
}
