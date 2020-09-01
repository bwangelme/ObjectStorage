package objects

import (
	"fmt"
	"io"
	"net/http"
)

// GetStream 读取文件的数据流
type GetStream struct {
	reader io.Reader
}

func newGetStream(url string) (*GetStream, error) {
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("data server return http code %d", r.StatusCode)
	}
	return &GetStream{r.Body}, nil
}

// NewGetStream 新建读取文件数据流
func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, object)
	}
	url := fmt.Sprintf("http://%s/objects/%s", server, object)
	return newGetStream(url)
}

// Read 数据流读取文件
func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
