package main

import (
	"flag"
	"fmt"

	"github.com/alimy/tl-explorer/internal"
)

var (
	host string
	port uint
)

func init() {
	flag.StringVar(&host, "host", "", "listening host")
	flag.UintVar(&port, "port", 8080, "listening port")
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", host, port)
	if host == "" {
		host = "localhost"
	}
	fmt.Printf("tl-explorer listening in [%s]. Please open http://%s:%d in browser to enjoy yourself.\n", addr, host, port)

	internal.StartServe(addr)
}
