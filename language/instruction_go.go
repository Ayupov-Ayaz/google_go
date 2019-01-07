package language

import (
	"fmt"
	"sync"
	"time"
)

//Горутина — это функция, которая может работать параллельно с другими функциями.
// Для создания горутины используется ключевое слово go, за которым следует вызов функции.
var printFormat = "| go, WaitGroup | Вывод потока: №% d : %d \n"
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

func TestGo2(count int, wg *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		defer wg.Done()

		wg.Add(1)
		go ts(i)
	}
}

func ts(j int){
	for i := 0; i < 10; i++ {
		fmt.Printf(printFormat, j + 1, i)
	}
}