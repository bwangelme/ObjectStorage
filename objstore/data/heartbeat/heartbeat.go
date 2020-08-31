package heartbeat

import (
	"time"

	"github.com/bwangelme/ObjectStorage/conf"
	"github.com/bwangelme/ObjectStorage/rabbitmq"
)

//StartHeartBeat 每隔5秒向 rabbitmq 发送一次心跳消息，供接口节点检测
func StartHeartBeat() {
	q := rabbitmq.New(conf.RabbitMQServer)
	defer q.Close()
	for {
		q.Publish(conf.ApiServersExchange, conf.ListenAddress)
		time.Sleep(5 * time.Second)
	}
}
