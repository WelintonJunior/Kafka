package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/IBM/sarama"
)

var (
	Brokers = []string{"localhost:9092"}
	Topic   = "chanMsg"
)

func main() {
	producer, err := getProducer()
	if err != nil {
		fmt.Println("Erro ao obter o producer:", err)
		return
	}
	defer producer.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite sua mensagem (Digite 'exit' para sair):")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		text = strings.TrimSpace(text)
		if strings.ToLower(text) == "exit" {
			break
		}

		msg := getMessage(text)

		_, _, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Erro ao mandar mensagem")
		}
	}
}

func getProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(Brokers, config)
}

func getMessage(msg string) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic:     Topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msg),
	}
}
