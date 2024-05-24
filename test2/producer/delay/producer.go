package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// 延迟消息和普通的发送区别 就是在需要发送的消息上，用下面的代码设置发送的级别即可
// message.WithDelayTimeLevel(3)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2), //指定重试次数
	)
	if err != nil {
		panic(err)
	}
	if err = p.Start(); err != nil {
		panic("启动producer失败")
	}
	topic := "test"
	// 构建一个消息
	message := primitive.NewMessage(topic, []byte("this is a delay message!"))
	// 给message设置延迟级别
	message.WithDelayTimeLevel(3)
	res, err := p.SendSync(context.Background(), message)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	if err = p.Shutdown(); err != nil {
		panic("关闭producer失败")
	}
}
