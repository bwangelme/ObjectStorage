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

func main() {
	flag.Parse()

	if conf.IsDataNode {
		go heartbeat.StartHeartBeat()
		go locate.StartLocate()
		http.HandleFunc("/objects/", objects.Handler)
		log.Fatalln(http.ListenAndServe(conf.ListenAddress, nil))
	} else {
		go aheartbeat.ListenHeartBeat()
		http.HandleFunc("/locate/", alocate.Handler)
		http.HandleFunc("/data/nodes", aheartbeat.Handler)
		log.Fatalln(http.ListenAndServe(conf.ListenAddress, nil))
	}
}
