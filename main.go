// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

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
	fmt.Printf("Listening in [%s]. Please open http://%s:%d in browser to enjoy yourself.\n", addr, host, port)

	internal.StartServe(addr)
}
