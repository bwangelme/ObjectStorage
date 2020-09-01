package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// 判断 rabbitmq server 是否可用

func main() {
	rabbitMQServer := os.Getenv("RABBITMQ_SERVER")
	conn, e := amqp.Dial(rabbitMQServer)
	if e != nil {
		log.Fatalln("Dial:", e)
	}

	ch, e := conn.Channel()
	if e != nil {
		log.Fatalln("Channel:", e)
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
		log.Fatalln("queue declare:", e)
	}
	fmt.Println("Declare Queue", q.Name)

	str, e := json.Marshal(map[string]string{
		"Hello": "World",
	})
	e = ch.Publish("", q.Name, false, false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    str,
		})
	if e != nil {
		log.Fatalln("publish:", e)
	}

	e = ch.Close()
	if e != nil {
		log.Fatalln("close channel:", e)
	}
}
