package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/devkhatri523/ecom-go/go-service/logger"
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	"strings"
)

type ConnectionDetail struct {
	Brokers  []string
	UseSSL   bool
	ClientId string
}

type ProducerConnectionDetail struct {
	ConnectionDetail ConnectionDetail
}
type ConsumerConnectionDetail struct {
	ConnectionDetail ConnectionDetail
	GroupName        string
	AutoCommit       bool
}

type Client interface {
	Start()
	Stop()
	Get() interface{}
	Subscribe(topic string) error
}

type ProducerClient struct {
	ProducerConnectionDetail ProducerConnectionDetail
	producer                 *kafka.Producer
}

type ConsumerClient struct {
	ConsumerConnectionDetail ConsumerConnectionDetail
	consumer                 *kafka.Consumer
}

func (p *ProducerClient) Start() {
	configMap := buildKafkaConfig(p.ProducerConnectionDetail.ConnectionDetail)
	strConfigMap := getJsonConfig(configMap)
	logger.Default().Debugf("Connecting to kafka producer with config : %s", strConfigMap)
	var err error
	producer, err := kafka.NewProducer(&configMap)
	if err != nil {
		logger.Default().Errorf("Error while connecting to kafka producer with config : %s  error : %s", strConfigMap, err)
	} else {
		logger.Default().Debugf("Connected to kafka producer with config : %s", strConfigMap)
	}
	p.producer = producer
}

func (p *ProducerClient) Subscribe(topic string) error {
	panic("Subscription of topic not required on producer client")
}
func (p *ProducerClient) Get() interface{} {
	return p.producer
}
func (p *ProducerClient) Stop() {
	if p.producer != nil {
		p.producer.Close()
		logger.Default().Debug("Kafka producer client stopped.")
	}
}

func (c *ConsumerClient) Start() {
	configMap := buildKafkaConfig(c.ConsumerConnectionDetail.ConnectionDetail)
	configMap.SetKey("group.id", c.ConsumerConnectionDetail.GroupName)
	strConfigMap := getJsonConfig(configMap)
	logger.Default().Debugf("Connecting to kafka consumer with config : %s", strConfigMap)
	consumer, err := kafka.NewConsumer(&configMap)
	if err != nil {
		logger.Default().Errorf("Error while connecting to kafka consumer with config : %s  error : %s", strConfigMap, err)
	} else {
		logger.Default().Debugf("Connected to kafka consumer with config : %s", strConfigMap)
	}
	c.consumer = consumer
}

func (c *ConsumerClient) Subscribe(topic string) error {
	return c.consumer.Subscribe(topic, nil)
}
func (c *ConsumerClient) Get() interface{} {
	return c.consumer
}
func (c *ConsumerClient) Stop() {
	if c.consumer != nil {
		err := c.consumer.Close()
		if err != nil {
			logger.Default().Errorf("Error while closing kafka consumer client error : %s", err)
		} else {
			logger.Default().Debug("Kafka consumer client stopped.")
		}
	}
}

func getJsonConfig(configMap kafka.ConfigMap) string {
	j, err := json.Marshal(configMap)
	if err == nil {
		return string(j)
	} else {
		return ""
	}
}

func buildKafkaConfig(detail ConnectionDetail) kafka.ConfigMap {
	clientId := detail.ClientId
	if utils.IsBlank(clientId) {
		clientId = utils.GenerateUUID()
	}
	if len(detail.Brokers) <= 0 {
		panic("Brokers list is empty.")
	}
	brokers := strings.Join(detail.Brokers[:], ",")
	configMap := kafka.ConfigMap{}
	configMap.SetKey("bootstrap.servers", brokers)
	configMap.SetKey("client.id", clientId)
	configMap.SetKey("acks", "all")
	return configMap

}
