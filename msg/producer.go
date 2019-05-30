package msg

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func Produ(strMessage string)  {
	syncProducer(strMessage)
}
func syncProducer(strMsg string) {
	fmt.Printf("this is a testing file \n")

	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer(Address, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	//构建发送的消息，
	msg := &sarama.ProducerMessage {
		Topic:    topicName,
		Partition: int32(10),
		Key:        sarama.StringEncoder("key"),
	}
	value := strMsg
	msg.Value = sarama.StringEncoder(value)
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		fmt.Printf("producer.go-->syncProducer SendMassage err: %s\n", err)
		panic(err)
	}
	fmt.Sprintf("%s_partition_%d\n", partition, offset)
}
