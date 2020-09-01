package objects

import (
	"fmt"
	"io"
	"net/http"
)

// PutStream 数据节点的写入数据流
type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

// NewPutStream new
func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		url := fmt.Sprintf("http://%s/objects/%s", server, object)
		request, _ := http.NewRequest("PUT", url, reader)
		client := http.Client{}
		// 请求会在 Do 这一步阻塞，直到 reader 读取到 io.EOF
		r, e := client.Do(request)
		// 状态码不是200 也认为是一种错误
		if e == nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{writer, c}
}

// Write 数据流写入文件数据
func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

// Close 关闭数据流
func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}
