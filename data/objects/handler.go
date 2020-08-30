package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/bwangelme/ObjectStorage/conf"
)

func get(w http.ResponseWriter, r *http.Request) {
	objectPath := path.Join(conf.StorageRoot, "objects")
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	f, err := os.Open(path.Join(objectPath, name))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func put(w http.ResponseWriter, r *http.Request) {
	objectPath := path.Join(conf.StorageRoot, "objects")
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	f, err := os.Create(path.Join(objectPath, name))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)

}

//Handler 处理文件的获取和添加请求
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodPut {
		put(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
