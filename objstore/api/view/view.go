package view

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bwangelme/ObjectStorage/api/heartbeat"
	"github.com/bwangelme/ObjectStorage/api/locate"
	"github.com/bwangelme/ObjectStorage/api/objects"
)

// LocateHandler 定位文件的接口
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

// AllNodesHandler 获取所有数据节点的接口
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

// ObjectsHandler 文件读写的接口
func ObjectsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		objects.Put(w, r)
		return
	case http.MethodGet:
		objects.Get(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
