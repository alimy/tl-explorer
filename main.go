// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/alimy/tl-explorer/assets"
)

var (
	host        string
	port        uint
	showVersion bool

	version = "v0.2.0"
)

func init() {
	flag.StringVar(&host, "host", "", "listening host")
	flag.UintVar(&port, "port", 8080, "listening port")
	flag.BoolVar(&showVersion, "v", false, "show version")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println(version)
		return
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	if host == "" {
		host = "localhost"
	}
	fmt.Printf("listening in [%s]. Please open http://%s:%d in browser to enjoy yourself.\n", addr, host, port)

	http.Handle("/", http.FileServer(assets.NewFileSystem()))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
