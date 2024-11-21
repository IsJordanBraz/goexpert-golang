package main

import "time"

type Message struct {
	id  int
	Msg string
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	// RabbitMQ
	go func() {
		msg1 := Message{id: 1, Msg: "RabbitMQ"}
		ch1 <- msg1
		time.Sleep(time.Second)
	}()

	// Kafka
	go func() {
		msg2 := Message{id: 2, Msg: "Kafka"}
		ch2 <- msg2
		time.Sleep(time.Second * 2)
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			println(msg1.Msg)
		case msg2 := <-ch2:
			println(msg2.Msg)
		case <-time.After(time.Second * 3):
			println("timeout")
		default:
			println("nothing ready")
		}
	}

}
