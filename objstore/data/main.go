package data

import (
	"log"
	"net/http"

	"github.com/bwangelme/ObjectStorage/conf"
	"github.com/bwangelme/ObjectStorage/data/heartbeat"
	"github.com/bwangelme/ObjectStorage/data/locate"
	"github.com/bwangelme/ObjectStorage/data/objects"
)

// Main date 功能入口
func Main() {
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatalln(http.ListenAndServe(conf.ListenAddress, nil))
}
