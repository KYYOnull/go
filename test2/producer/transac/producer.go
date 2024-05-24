package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// 发送事务消息需要写TransactionListener接口的方法
// 指明事务执行成功和回调的具体操作，接口如下
// type TransactionListener interface {
// 	//  When send transactional prepare(half) message succeed, this method will be invoked to execute local transaction.
// 	ExecuteLocalTransaction(*Message) LocalTransactionState

// 	// When no response to prepare(half) message. broker will send check message to check the transaction status, and this
// 	// method will be invoked to get local transaction status.
// 	CheckLocalTransaction(*MessageExt) LocalTransactionState
// }

type DemoListener struct {
}

func (dl *DemoListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second * 3)
	fmt.Println("执行本地逻辑失败")
	//本地执行逻辑无缘无故失败 代码异常 宕机
	return primitive.UnknowState
}

func (dl *DemoListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("rocketmq的消息回查")
	return primitive.CommitMessageState
}

func main() {
	p, _ := rocketmq.NewTransactionProducer(
		&DemoListener{},
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(1),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
		os.Exit(1)
	}

	res, err := p.SendMessageInTransaction(context.Background(),
		primitive.NewMessage("TopicTest5", []byte("Hello RocketMQ again ")))

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	time.Sleep(5 * time.Minute)
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
