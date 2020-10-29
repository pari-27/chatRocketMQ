package main

import (
	"chatRocketMQ/wrapper"

	"fmt"
)

func main() {
	// pMsg := ""

	rMsg := make(chan string, 100)

	go wrapper.ReceiveMsg(rMsg)

	go wrapper.SendMsg("new hello555555")

	fmt.Println(<-rMsg)

	// for {
	// 	fmt.Println("enter message:")
	// 	fmt.Scan(&pMsg)
	// 	wrapper.SendMsg(pMsg)

	// 	// fmt.Printf("%v", rMsg)
	// }
}
