package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bwangelme/ObjectStorage/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatalln(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
