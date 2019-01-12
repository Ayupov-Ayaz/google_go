package language

import (
	"fmt"
	"time"
)
// Событие которое будет происходить каждую секунду. howLongSec - сколько секунд
func StartTicker(howLongSec int) {
	ticker := time.NewTicker(time.Second)
	t := 0
	for tickTime := range ticker.C { // Тикер работает через свой канал
		t++
		fmt.Printf("step %d, time: %v \n", t, tickTime)
		if t >= howLongSec {
			ticker.Stop()
			break
		}
	}
	fmt.Println("total ", t)
}