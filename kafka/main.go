package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
)

// CustomPartitioner 是一个自定义的分区器
type CustomPartitioner struct{}

func (p *CustomPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	// 自定义分区逻辑，例如根据消息的Key来分区
	return int32(message.Key.Length() % int(numPartitions)), nil
}

func (p *CustomPartitioner) RequiresConsistency() bool {
	return true
}

// ConsumerGroupHandler 实现 sarama.ConsumerGroupHandler 接口
type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	// 在重新平衡开始时调用，可以用于执行一些初始化操作
	fmt.Println("Setup: Consumer group is rebalancing")
	return nil
}

func (ConsumerGroupHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	// 在重新平衡完成时调用，可以用于执行一些清理操作
	fmt.Println("Cleanup: Consumer group rebalancing finished")
	return nil
}

func (ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(msg.Value), msg.Timestamp, msg.Topic)
		sess.MarkMessage(msg, "")
		// 手动提交偏移量
		sess.Commit()
	}
	return nil
}

func main() {
	// 设置日志输出到标准输出
	log.SetOutput(os.Stdout)

	// 启动生产者
	go startProducer()

	// 启动消费者
	startConsumer()
}

func startProducer() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = func(topic string) sarama.Partitioner {
		return &CustomPartitioner{}
	}
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	defer producer.Close()

	for i := 0; i < 10; i++ {
		msg := &sarama.ProducerMessage{
			Topic: "test_topic",
			Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("value-%d", i)),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}

		fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		time.Sleep(1 * time.Second)
	}
}

func startConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Fetch.Min = 1
	config.Consumer.MaxWaitTime = 100 * time.Millisecond

	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "test_group", config)
	if err != nil {
		log.Fatalf("Failed to create Sarama consumer group: %v", err)
	}
	defer consumerGroup.Close()

	handler := ConsumerGroupHandler{}

	for {
		err := consumerGroup.Consume(context.Background(), []string{"test_topic"}, handler)
		if err != nil {
			log.Fatalf("Error consuming messages: %v", err)
		}
	}
}
