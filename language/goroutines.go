package language

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const(
	goroutinesCount = 5
	iterationsCount = 5
)
var printFormatGoroutines = "%s%s th%d %s%s%d%s \n"
// Для задержки выполнения, что бы главный поток (main) ждал другие
//Горутина — это функция, которая может работать параллельно с другими функциями.
// Для создания горутины используется ключевое слово go, за которым следует вызов функции.

func RunGoroutines() {
	var wg = &sync.WaitGroup{} // всегда по ссылке! нет смысла копировать его
	// workers
	for i := 0; i < goroutinesCount; i++ {
		// выставляем в очередь нашу горутину в WaitGroup
		wg.Add(1)
		// создаем горутину
		go doSomeWork(i, wg)
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
}

func doSomeWork(in int, wg *sync.WaitGroup) {
	for j := 0; j < iterationsCount; j++ {
		fmt.Printf(formatWork(in, j))
		// переключиться на другую горутину
		runtime.Gosched()
	}
	wg.Done()
}

func formatWork(in, j int) string {
	return fmt.Sprintf(printFormatGoroutines, strings.Repeat(" ",in), "∎", in,
		strings.Repeat(" ", j),"j", j, strings.Repeat("∎", j))
}
