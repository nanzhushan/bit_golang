package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	var address = []string{"192.168.100.12:9092"}
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer p.Close()
	topic := "xxoo"
	value := "aa"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", value, err)
	} else {
		fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
	}
	time.Sleep(2 * time.Second)

}

