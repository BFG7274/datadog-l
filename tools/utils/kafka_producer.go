package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/Shopify/sarama"
)

var (
	producer sarama.SyncProducer
)

var conf Conf

type Conf struct {
	Brokers string `yaml:"brokers"`
	Group   string `yaml:"group"`
	Topics  string `yaml:"topics"`
	Enable  bool
}

func init() {
	enableStr := strings.ToLower(os.Getenv("AIOPS_KAFKA"))
	if enableStr == "" || enableStr == "false" {
		return
	}
	conf.Enable = true
	conf.Brokers = os.Getenv("AIOPS_KAFKA_BROKERS")

	initProducer()
}

func initProducer() {

	var err error
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 3
	config.Producer.Return.Successes = true
	brokers := strings.Split(conf.Brokers, ",")
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Printf("kafka生产者初始化失败 -> %v \n", err)
		panic(err)
	}
	fmt.Println("kafka生产者初始化成功")
}

func SendKafka(topic string, msg []byte) {
	if !conf.Enable {
		return
	}
	//生产消息
	_, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	})
	if err != nil {
		log.Warn("send kafka failed, err:", err)
		return
	}
	log.Warnf("send kafka success, offset: %v\n", offset)
}
