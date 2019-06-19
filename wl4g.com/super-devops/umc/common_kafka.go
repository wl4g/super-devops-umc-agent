package main

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

var producer sarama.SyncProducer

func initKafka()  {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	var err error
	producer, err = sarama.NewSyncProducer([]string{conf.KafkaConf.Url}, config)
	if err != nil {
		panic(err)
	}
}

func send(text string)  {

	//构建发送的消息，
	msg := &sarama.ProducerMessage {
		Topic:conf.KafkaConf.Topic,
		Value:sarama.ByteEncoder(text),
		Partition: int32(conf.KafkaConf.Partiations),//
		//Key:        sarama.StringEncoder("key"),//
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		MainLogger.Error("Send message Fail",zap.Error(err))
	}

	MainLogger.Info("Send message Success - ",
		zap.Int32("Partition", partition),
		zap.Int64("offset", offset))

}

