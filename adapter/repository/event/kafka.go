package event

import (
	"context"
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorEventRepositoryImpl struct {
	kp *kafka.Producer
}

func (author *AuthorEventRepositoryImpl) Publish(ctx context.Context, key, message []byte) (err error) {
	var (
		topics       = config.GetStringSlice("kafka.topics")
		topic        string
		deliveryChan = make(chan kafka.Event)
	)

	if len(topics) > 0 {
		topic = topics[0]
	}
	err = author.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
		Key:            key,
	}, deliveryChan)

	e := <-deliveryChan

	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return errors.New(fmt.Sprintf("Delivery failed: %v\n", m.TopicPartition.Error))
	}
	close(deliveryChan)
	return
}

func NewAuthorEventRepository(kp *kafka.Producer) repository.AuthorEventRepository {
	return &AuthorEventRepositoryImpl{
		kp: kp,
	}
}
