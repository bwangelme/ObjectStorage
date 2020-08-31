package locate

import (
	"log"
	"os"
	"path"
	"strconv"

	"github.com/bwangelme/ObjectStorage/conf"
	"github.com/bwangelme/ObjectStorage/rabbitmq"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(conf.RabbitMQServer)
	defer q.Close()
	q.Bind(conf.DataServersExchange)
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			log.Fatalln(e)
		}
		objPath := path.Join(conf.StorageRoot, "objects", object)
		if Locate(objPath) {
			q.Send(msg.ReplyTo, conf.ListenAddress)
		}
	}
}
