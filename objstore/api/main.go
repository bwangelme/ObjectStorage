package api

import (
	"log"
	"net/http"

	"github.com/bwangelme/ObjectStorage/api/heartbeat"
	"github.com/bwangelme/ObjectStorage/api/view"
	"github.com/bwangelme/ObjectStorage/conf"
)

func Main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/locate/", view.LocateHandler)
	http.HandleFunc("/objects/", view.ObjectsHandler)
	http.HandleFunc("/data/nodes", view.AllNodesHandler)
	log.Fatalln(http.ListenAndServe(conf.ListenAddress, nil))
}
