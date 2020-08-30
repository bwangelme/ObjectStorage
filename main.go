package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bwangelme/ObjectStorage/conf"
	"github.com/bwangelme/ObjectStorage/data/heartbeat"
	"github.com/bwangelme/ObjectStorage/data/locate"
	"github.com/bwangelme/ObjectStorage/data/objects"

	aheartbeat "github.com/bwangelme/ObjectStorage/api/heartbeat"
	alocate "github.com/bwangelme/ObjectStorage/api/locate"
)

var isData = flag.Bool("data", false, "Is Running Data Node")

func main() {
	flag.Parse()

	if *isData {
		go heartbeat.StartHeartBeat()
		go locate.StartLocate()
		http.HandleFunc("/objects/", objects.Handler)
		log.Fatalln(http.ListenAndServe(conf.ListenAddress, nil))
	} else {
		go aheartbeat.ListenHeartBeat()
		http.HandleFunc("/locate/", alocate.Handler)
		log.Fatalln(http.ListenAndServe(conf.ListenAddress, nil))
	}
}
