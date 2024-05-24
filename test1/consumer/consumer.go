package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	if err != nil {
		panic(err)
	}
	err = c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("subscribe callback: %+v \n", msg)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(err)
	}
	err = c.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = c.Shutdown()
		if err != nil {

			fmt.Printf("shutdown Consumer error: %s", err.Error())
		}
	}()
	<-(chan interface{})(nil)

}
