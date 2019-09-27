package main

import (
	"fmt"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func Consume(message chan string, closeChan chan byte) {
	brokers := []string{viper.GetString("kafka.brokerList")}
	topics := []string{viper.GetString("kafka.topic")}
	groupId := viper.GetString("kafka.groupId")

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	go func() {
		for err := range consumer.Errors() {
			log.Infof("Consumer Error: %s\n", err.Error())
		}
	}()

	go func() {
		for ntf := range consumer.Notifications() {
			log.Infof("Consumer Rebalanced: %+v", ntf)
		}
	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				//log.Infof("received msg: %s\n", msg.Value)
				message <- string(msg.Value)
				consumer.MarkOffset(msg, "")
			} else {
				log.Error("received kafka msg error", fmt.Errorf(""))
			}
		case <-closeChan:
			break
		}

	}

}
