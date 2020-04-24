package main

import "flag"

var (
	ip     *string = flag.String("ip", "127.0.0.1", "Define server and ip")
	server *bool   = flag.Bool("server", false, "Download server files")
)
