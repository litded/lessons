package main

import "fmt"

func PrintInfo(win chan string) {
	fmt.Println("Length of channel:", len(win), "Cap of channel:", cap(win))
}

func main() {
	window := make(chan string, 2)
	PrintInfo(window)
	window <- "TOSHIBA"
	PrintInfo(window)
	window <- "LG"
	PrintInfo(window)
	fmt.Println(<-window, <-window)
	PrintInfo(window)

}
