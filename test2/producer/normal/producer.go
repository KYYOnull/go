package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

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
	message := primitive.NewMessage(topic, []byte("hello world!"))
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
