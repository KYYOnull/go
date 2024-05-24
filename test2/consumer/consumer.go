package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	if err != nil {
		panic(err)
	}

	if err := c.Subscribe(
		"test",
		consumer.MessageSelector{},
		// 收到消息后的回调函数
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Printf("获取到值： %v \n", msgs[i])
			}
			return consumer.ConsumeSuccess, nil
		}); err != nil {
		fmt.Println(err)
	}
	err = c.Start()
	if err != nil {
		panic("启动consumer失败")
	}
	//不能让主goroutine退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
