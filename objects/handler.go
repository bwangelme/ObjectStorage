package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	objectPath := path.Join(os.Getenv("STORAGE_ROOT"), "objects")
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
	objectPath := path.Join(os.Getenv("STORAGE_ROOT"), "objects")
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
