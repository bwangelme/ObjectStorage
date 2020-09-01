package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bwangelme/ObjectStorage/api/heartbeat"
	"github.com/bwangelme/ObjectStorage/api/locate"
)

// Put 向数据节点写入文件
func Put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(c)
}

// Get 从数据节点获取文件
func Get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}

func putStream(object string) (*PutStream, error) {
	server := heartbeat.ChooseRandomDataNode()
	if server == "" {
		return nil, fmt.Errorf("cannot find any data server")
	}
	return NewPutStream(server, object), nil
}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return NewGetStream(server, object)
}