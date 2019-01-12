package language

import (
	"fmt"
	"sync"
)
var wgCh = sync.WaitGroup{}
var writeFormat string = "| chan | Writing in channel: %s\n"
var readeFormat string = "| chan | Reading from channel: %s\n"

// Канал может только записывать
func send(channel chan <- string, data string){

	for _, char := range data {
		charStr := string(char)
		fmt.Printf(writeFormat, charStr)
		channel <- charStr
	}
	//close(channel)
}

// Канал может только читать
func printer(channel <- chan string){
	for char := range channel {
		fmt.Printf(readeFormat, char)
	}
	wgCh.Done()
}

func TestChannel(){
	var channel = make(chan string) // создание канала
	var testStr = "Hello world!!!"
	go send(channel, testStr)
	go printer(channel)

}