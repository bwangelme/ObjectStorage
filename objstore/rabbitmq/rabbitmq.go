package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

//RabbitMQ RabbitMQ 客户顿
type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	exchange string
}

//New 创建 RabbitMQ 客户端
func New(serverAddr string) *RabbitMQ {
	conn, e := amqp.Dial(serverAddr)
	if e != nil {
		log.Fatal(e)
	}

	ch, e := conn.Channel()
	if e != nil {
		log.Fatalln(e)
	}

	q, e := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if e != nil {
		log.Fatalln(e)
	}

	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name

	return mq
}

//Close 关闭 RabbitMQ 客户端
func (q *RabbitMQ) Close() {
	// TODO: 是否也要将连接关闭
	q.channel.Close()
}

//Bind 将队列绑定到指定交换机上
// exchange 是 direct 类型，bindKey 默认是空
func (q *RabbitMQ) Bind(exchange string) {
	e := q.channel.QueueBind(
		q.Name, // Queue Name
		"",     // BindKey
		exchange,
		false,
		nil)
	if e != nil {
		log.Fatalln(e)
	}
	q.exchange = exchange
}

//Send 将消息发送到指定队列上，且要求对端也将返回消息写入到该队列中
//这里使用的是默认的交换机, `""`
func (q *RabbitMQ) Send(queue string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		log.Fatalln(e)
	}
	e = q.channel.Publish("", queue, false, false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    str,
		})
	if e != nil {
		log.Fatalln(e)
	}
}

//Publish 向某个 Exchange 发送消息
// Exchange 是 direct 类型，所以发送的消息是一对多的关系
func (q *RabbitMQ) Publish(exchange string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		log.Fatalln(e)
	}

	e = q.channel.Publish(exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    str,
		})
	if e != nil {
		log.Fatalln(e)
	}
}

//Consume 从队列中接收消息
func (q *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := q.channel.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if e != nil {
		log.Fatalln(e)
	}
	return c
}
