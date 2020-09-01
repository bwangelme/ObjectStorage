package locate

import (
	"strconv"
	"time"

	"github.com/bwangelme/ObjectStorage/conf"
	"github.com/bwangelme/ObjectStorage/rabbitmq"
)


func Locate(name string) string {
	q := rabbitmq.New(conf.RabbitMQServer)
	q.Publish(conf.DataServersExchange, name)
	c := q.Consume()
	go func() {
		//如果队列在一秒内未返回消息，则认为该对象不存在
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
