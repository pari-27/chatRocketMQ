package wrapper

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

//ReceiveMsg to get messages from the queue
func ReceiveMsg(rMsg chan string) {
	var message []*primitive.MessageExt

	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("new2TestGroup"),
		consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)

	if err != nil {
		fmt.Println("helllloooooo" + err.Error())
	}

	err = c.Subscribe("SelfTest3P", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		message = msgs
		fmt.Printf("received msg: %v \n", message)
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println("helllloooooo" + err.Error())
	}

	// fmt.Println("subscription done")
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	for m := range message {
		rMsg <- message[m].String()
	}
	time.Sleep(time.Hour)

	err = c.Shutdown()
	if err != nil {
		fmt.Printf("Shutdown Consumer error: %s", err.Error())
	}

	close(rMsg)

}
