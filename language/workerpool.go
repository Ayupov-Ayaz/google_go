package language

import (
	"fmt"
	"runtime"
	"time"
)
 /**
 	Worker pool
 	Довольно часто в архитектуре встречаются случаи, когда работа осуществляется через какие-то очереди, с которыми
 	работают какие-то офлайновые уже «разгребаторы».Довольно часто в архитектуре встречаются случаи,
 	когда работа осуществляется через какие-то очереди, с которыми работают какие-то офлайновые уже «разгребаторы».
 	Мы можем создать сразу несколько горутин, которые будут читать из какого-то канала.
 	И если нам приходит какая-то задача, мы просто будем писать ее в этот канал.
 	А там воркеры уже дальше сами разберутся
  */
func StartWorkerPool(workerCount int) {
	workerInput := make(chan string, 2)
	for i := 0; i < workerCount; i++ {
		go startWorker(workerInput, i)
	}
	months := []string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль",
					   "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"}

	// Распределение задач по воркерам
	for _, month := range months {
		workerInput <- month
	}
	close(workerInput) // Закрываем канал
	time.Sleep(time.Millisecond)
}

// Запуск воркеров которые в нашем случае просто выводят результат на экран.
func startWorker(in <- chan  string, workerNum int) {
	 for input := range in {
	 	fmt.Printf("Worker №%d sayd: %s \n", workerNum + 1 , input)
	 	runtime.Gosched() // переключаемся на другого воркера
	 }
	 fmt.Printf("=== worker №%d is done === \n", workerNum + 1)
}