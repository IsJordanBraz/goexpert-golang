package main

// Thread 1
func main() {
	forever := make(chan bool)

	go func() {
		forever <- true
	}()

	<-forever
}
