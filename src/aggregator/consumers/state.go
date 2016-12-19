package consumers

import (
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	saramaC "github.com/d33d33/sarama-cluster"
	"github.com/spf13/viper"

	"github.com/runabove/metronome/src/metronome/kafka"
	"github.com/runabove/metronome/src/metronome/models"
	"github.com/runabove/metronome/src/metronome/redis"
)

// StateConsumer consumed states messages from Kafka to maintain the state database.
type StateConsumer struct {
	consumer *saramaC.Consumer
}

// NewStateConsumer returns a new state consumer.
func NewStateConsumer() (*StateConsumer, error) {
	brokers := viper.GetStringSlice("kafka.brokers")

	config := saramaC.NewConfig()
	config.Config = *kafka.NewConfig()
	config.ClientID = "metronome-aggregator"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := saramaC.NewConsumer(brokers, kafka.GroupAggregators(), []string{kafka.TopicStates()}, config)
	if err != nil {
		return nil, err
	}

	sc := &StateConsumer{
		consumer,
	}

	go func() {
		for {
			select {
			case msg, ok := <-consumer.Messages():
				if !ok { // shuting down
					return
				}
				sc.handleMsg(msg)
			}
		}
	}()

	return sc, nil
}

// Close the consumer.
func (sc *StateConsumer) Close() error {
	return sc.consumer.Close()
}

// Handle message from Kafka.
// Apply updates to the database.
func (sc *StateConsumer) handleMsg(msg *sarama.ConsumerMessage) {
	var s models.State
	if err := s.FromKafka(msg); err != nil {
		log.Error(err)
		return
	}

	log.Infof("UPDATE state: %s", s.TaskGUID)

	if err := redis.DB().HSet(s.UserID, s.TaskGUID, s.ToJSON()).Err(); err != nil {
		panic(err)
	}
	if err := redis.DB().PublishTopic(s.UserID, "state", s.ToJSON()).Err(); err != nil {
		log.Error(err)
	}
}
