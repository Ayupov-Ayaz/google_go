package language

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/**
Context - нужен в основном для отмены асинхронных операций.

 */
func StartContext() {
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