package msg

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

/*
	Global variable
 */
//var Address = []string{"10.110.25.114:9092", "10.110.25.113:9092", "10.110.25.108:9092"}
var (
	topicName = "monitor-metrics"
	wg sync.WaitGroup
	)

func Runconsumer()  {
	fmt.Printf("this is a consumer \n")
	getconsumer()
}

func getconsumer()  {
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer(Address, nil)
	if err != nil {
		panic(err)
	}
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions(topicName)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition(topicName, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:\n%s\n", msg.Topic,msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
