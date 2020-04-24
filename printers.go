package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type DownloadPrinter struct {
	Total uint64
}

func (printer *DownloadPrinter) Write(file []byte) (int, error) {
	n := len(file)

	printer.Total += uint64(n)

	fmt.Printf("\rDownloading... %s ", humanize.Bytes(printer.Total))

	return n, nil
}
