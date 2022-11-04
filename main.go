package main

import (
	"awesomeProject/02-messageAck/consumer"
	"fmt"
)

// 测试手动应答，当一个消费者挂机后，消息重新入队
func main() {
	forever := make(chan interface{})
	fmt.Println("Please select a consumer to start: 1 fast consumer, 2 slow consumer")
	var mode int
	fmt.Scanf("%d", &mode)
	if mode == 1 {
		go consumer.FastConsumer()
	} else if mode == 2 {
		go consumer.SlowConsumer()
	} else {
		fmt.Println("illegal selection!")
	}
	<-forever
}
