package main

import (
	"fmt"
	"time"
)

func printMesage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second * 1)
	}
	channel <- "Done!"
}

func main() {
	var channel chan string
	go printMesage("FM rocks", channel)
	response := <- channel

	fmt.Println(response)
}
