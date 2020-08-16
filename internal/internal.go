// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

func StartServe(addr string) {
	box := packr.New("schemaTL", "../assets/schema.tl")

	http.Handle("/", http.FileServer(box))
	log.Fatal(http.ListenAndServe(addr, nil))
}
