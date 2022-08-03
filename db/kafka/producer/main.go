package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func main() {
	// to create topics when auto.create.topics.enable='false'
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		fmt.Println("dial:", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		fmt.Println("Controller:", err)
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		fmt.Println("Dial2:", err)
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
