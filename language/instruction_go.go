package language

import (
	"fmt"
	"sync"
	"time"
)

func TestGo(millisecond time.Duration,wg *sync.WaitGroup){
	// Ставим в очередь
	wg.Add(1)
	// запускаем функцию в отдельном потоке
	go func() {
		// нужно обязательно закрыть этот поток
		defer wg.Done()
		duration := millisecond * time.Millisecond
		time.Sleep(duration) // инициируем какие то расчеты на что тратится время
		fmt.Println(duration)
	}()
}