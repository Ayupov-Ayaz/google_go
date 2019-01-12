package language

import (
	"fmt"
	"time"
)

func StartAfterFunc(howOftenSec int) {
	time.AfterFunc(time.Second * time.Duration(howOftenSec), sayHello)
}

func sayHello() {
	fmt.Printf("Hello world!!!! \n")
}