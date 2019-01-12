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

/**
	Бесконечный тикер
	Если завершить цикл и не выйти из программы то тикер будет продолжать работать
	Его нужно использовать тогда, когда вы точно не собираетесь его останавливать: Мониторинг, сбор и отсылка каких либо
 	данных и т.д
 */
func InfinityTicker(howLongSec int) {
	ticker := time.Tick(time.Second)
	t := 0
	// сразу возвращает канал с которого можно считать
	for chanel := range ticker {
		t++
		fmt.Printf("step %d, time: %v \n", t, chanel)
		if t >= howLongSec {
			break
		}
	}
}