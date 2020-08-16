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
