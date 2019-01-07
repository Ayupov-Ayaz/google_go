package language

import (
	"fmt"
	"time"
)

func TestSelect(data string){
	printFormat  := "| select | from chanel %s -> data: %s \n"

	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		for _, char := range data{
			c1 <- string(char)
		}
	}()

	go func() {
		for _, char := range data{
			c2 <- string(char)
		}
	}()

	go func() {
		for{
			select {
				case msg1 := <- c1:
					fmt.Printf(printFormat, "c1", msg1)
				case msg2 := <- c2:
					fmt.Printf(printFormat, "c2", msg2)
					// обычно select используется для таймеров
				case <- time.After(time.Second):
					fmt.Println("timeout")
				default:
					fmt.Printf("| select | nothing ready =( \n")
			}
		}

	}()
}