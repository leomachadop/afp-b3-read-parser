package processor

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func SendToKafka(histologically []byte, wg sync.WaitGroup) error {

	topic := "b3-historical-quotations"

	brokersUrl := []string{"192.168.56.1:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		wg.Done()
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(histologically),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("Error send message to kafka %s", err)
		wg.Done()
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	wg.Done()

	return nil
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
