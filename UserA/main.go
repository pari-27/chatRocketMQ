package main

import (
	"chatRocketMQ/wrapper"

	"fmt"
)

func main() {
	// pMsg := ""

	rMsg := make(chan string, 100)
	var newMsg string
	go wrapper.ReceiveMsg(rMsg)

	fmt.Println("enter new msg:")
	fmt.Scanln(&newMsg)

	go wrapper.SendMsg(newMsg)
	fmt.Println(<-rMsg)
	// for {
	// 	fmt.Println("enter message:")
	// 	fmt.Scan(&pMsg)
	// 	wrapper.SendMsg(pMsg)

	// 	// fmt.Printf("%v", rMsg)
	// }
}
