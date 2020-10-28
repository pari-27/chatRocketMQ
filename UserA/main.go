package main

import (
	"chatRocketMQ/wrapper"
	"fmt"
	"time"
)

func main() {
	// pMsg := ""

	rMsg := make(chan string, 100)

	wrapper.ReceiveMsg(rMsg)
	wrapper.SendMsg("new hello")

	fmt.Println("hiiii555555" + <-rMsg)

	time.Sleep(time.Hour)

	// for {
	// 	fmt.Println("enter message:")
	// 	fmt.Scan(&pMsg)
	// 	wrapper.SendMsg(pMsg)

	// 	// fmt.Printf("%v", rMsg)
	// }
}
