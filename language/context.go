package language

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/**
Context - нужен в основном для отмены асинхронных операций.
в ContextCancel сыметирована ситуация когда нам нужно получить первый результат из любого канала
 */
func StartContextWithCancel() {
	// Создаем контекст с функцией отмены. контекс background это просто как родительский тип
	ctx, cancelFunc := context.WithCancel(context.Background()) // нам возвращается сам контекст и функция отмены
	c := make(chan int, 1)

	// создаем 10 воркеров
	for i := 0; i < 10; i++ {
		go makeWorker(ctx, i, c)
	}
	// дожидаемся первого результат и останавливаем
	fondBy := <-c
	fmt.Println("result fond by", fondBy)
	// функция обрабатывается если выполняется метод контекста Done()
	cancelFunc()
	time.Sleep(time.Second)
}

func makeWorker(ctx context.Context, workerNum int, out chan <- int) {
	waitTime := time.Duration(rand.Intn(100) + 10) * time.Millisecond
	fmt.Printf("Worker №%d sleep %d milliseconds \n", workerNum, waitTime)
	select {
	case <- ctx.Done(): //если придет сигнал, что нужно закругляться, мы просто выйдем из функции
		return
	case <- time.After(waitTime): // имитируем наш sleep
		fmt.Println("Worker №", workerNum, "done")
		out <- workerNum //
	}
}

/**
 Завершение работы по истечению какого то времени
 */
func StartContextWithTimeout(millisec int) {
	workTime := time.Duration(millisec)  * time.Millisecond
	ctx, _ := context.WithTimeout(context.Background(), workTime)
	c := make(chan int, 1)

	// создаем воркеры
	for i := 0; i < 10; i++ {
		go makeWorker(ctx, i, c)
	}
	// сколько раз будут получены результаты
	totalFound := 0

	LOOP:
		for {
			select {
			case <- ctx.Done():
				break LOOP
			case foundBy := <- c:
				totalFound++
				fmt.Println("result fond by ", foundBy)
			}
		}
	fmt.Println("total found ", totalFound)
	time.Sleep(time.Second)
}