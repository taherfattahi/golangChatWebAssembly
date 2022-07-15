package main

import (
	"log"
	"net/http"
)

const (
	AddSrv       = ":8080"
	TemplatesDir = "../app/html/"
)

func main() {
	fileSrv := http.FileServer(http.Dir(TemplatesDir))

	err := http.ListenAndServe(AddSrv, fileSrv)

	if err != nil {
		log.Fatalln(err)
	}
}
