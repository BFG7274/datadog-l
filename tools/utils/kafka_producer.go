package utils

import (
	"fmt"
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/Shopify/sarama"
)

var (
	producer sarama.SyncProducer
	brokers  = []string{"192.168.0.105:9092", "192.168.0.139:9092", "192.168.0.229:9092"}
)

func init() {
	initProducer()
}

func initProducer() {
	var err error
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 3
	config.Producer.Return.Successes = true
	brokers := brokers
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Printf("kafka生产者初始化失败 -> %v \n", err)
		panic(err)
	}
	fmt.Println("kafka生产者初始化成功")
}

func SendKafka(topic string, msg []byte) {
	//生产消息
	_, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	})
	if err != nil {
		log.Warn("send kafka failed, err:", err)
		return
	}
	log.Warn("send kafka success, offset: %v\n", offset)
}