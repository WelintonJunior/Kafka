package main

import (
	"fmt"
	db "kafka-consumer/database"

	"github.com/IBM/sarama"
)

func main() {
	db.InitDB()
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	partitionConsumer, _ := consumer.ConsumePartition("chanMsg", 0, sarama.OffsetNewest)

	channel := partitionConsumer.Messages()

	for {
		msg := <-channel
		err := sendMsgToDatabase(msg.Value)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(msg.Value))
		}
	}
}

func sendMsgToDatabase(msg []byte) error {
	stringMsg := string(msg)
	query := "insert into tblmessages (msg) values($1)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(stringMsg)

	if err != nil {
		return err
	}

	return nil
}
