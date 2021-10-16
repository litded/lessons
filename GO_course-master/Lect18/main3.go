package main

func main() {
	a := make(chan int)
	data := <-a // read from channel
	a <- data   //write to channel
}
