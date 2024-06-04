package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	partitionConsumer, _ := consumer.ConsumePartition("chanMsg", 0, sarama.OffsetNewest)

	channel := partitionConsumer.Messages()

	for {
		msg := <-channel
		fmt.Println(string(msg.Value))
	}
}
