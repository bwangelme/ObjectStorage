package view

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bwangelme/ObjectStorage/api/heartbeat"
	"github.com/bwangelme/ObjectStorage/api/locate"
	"github.com/bwangelme/ObjectStorage/api/objects"
)

func LocateHandler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info := locate.Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}

func AllNodesHandler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	nodes := heartbeat.AllNodes()
	b, _ := json.Marshal(map[string]interface{}{
		"nodes": nodes,
	})
	w.Write(b)
}

func ObjectsHandler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		objects.Put(w, r)
		return
	}
	// TODO: get 方法
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
