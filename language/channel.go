package language

import (
	"fmt"
)

var writeFormat string = "| chan | Writing in channel: %s\n"
var readeFormat string = "| chan | Reading from channel: %s\n"

// Канал может только записывать
func send(channel chan <- string, data string){

	for _, char := range data {
		charStr := string(char)
		fmt.Printf(writeFormat, charStr)
		channel <- charStr
	}
}

// Канал может только читать
func printer(channel <- chan string){
	for char := range channel {
		fmt.Printf(readeFormat, char)
	}
}

func TestChannel(data string){
	var channel = make(chan string) // создание канала
	go send(channel, data)
	go printer(channel)
}