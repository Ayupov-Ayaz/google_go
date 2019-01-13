package language

import (
	"fmt"
	"sync"
)
/**
	Мьютекс большой жирный примитив для синхронизации больших кусков программы или алгоритмов. Он реализован на основе
	примитива sync.Atomic . Для операций где используется счетчик можно использовать Atomic
 */
func StartRaceConditionWithMutex() {
	var counters = map[int]int{}
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				// залочили данные, что бы никто другой не мог ими пользоваться
				mu.Lock()
				counters[th * 10 +j]++ // fatal error, можно отследить если запустить программу с ключём -race
				// разлочили
				mu.Unlock()
			}
		}(counters, i, mu)
	}
	wg.Wait()
	fmt.Println(counters)
}