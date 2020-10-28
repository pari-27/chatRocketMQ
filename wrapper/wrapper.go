package wrapper

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

//ReceiveMsg to get messages from the queue
func ReceiveMsg(rMsg chan string) {

	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("firstTestGroup"),
		consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
		consumer.WithConsumerOrder(true),
	)

	err := c.Subscribe("SelfTest1P", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		fmt.Printf("received msg: %v\n", msgs)
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	rMsg <- "success"
	time.Sleep(time.Hour)

	err = c.Shutdown()
	if err != nil {
		fmt.Printf("Shutdown Consumer error: %s", err.Error())
	}
	close(rMsg)
}

//SendMsg to send a message to a queue
func SendMsg(Pmsg string) {

	p, _ := rocketmq.NewProducer(
		producer.WithGroupName("firstTestGroup"),
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	topic := "SelfTest1P"
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(Pmsg),
	}
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}

}
