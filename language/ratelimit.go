package language

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)


const (
	quotaLimit = 2
	quotaGoroutinesCount = 3
	quotaIterationsCount = 6
)

/**
	Довольно часто бывает так, что нам нужно каким-либо образом затормозить нашу программу, например, в зависимости
	от утилизации процессора, либо дисковой подсистемы, нам нужно уменьшить нагрузки на эти части.
	В Go мы можем организовать это, используя буферизированные каналы. */
func StartRiteLimit() {
	 wg := &sync.WaitGroup{}
	 // Создаем канал с лимитом квоты. По сути это буферизированный канал где размер буфера и является нашей квотой
	 quotaCh := make(chan struct{}, quotaLimit)

	 for i := 0; i < quotaGoroutinesCount; i++ {
	 	wg.Add(1)
	 	go startWorkerWithQuota(wg, i, quotaCh)
	 }
	 time.Sleep(time.Millisecond)
	 wg.Wait()
}

func startWorkerWithQuota(wg *sync.WaitGroup, workerNum int, quotaCh chan struct{}) {
	 /** прежде чем начать работу, воркер попытается получить слот на эту работу то есть, мы пытаемся положить пустую
	 	структуру в канал. Если буфер канала уже заполнен, значит 2 воркера уже работают, то есть там уже есть 2 пустые
	    структуры. Мы блокируемся до тех пор пока не освободится место  */
	quotaCh <- struct{}{}
	defer wg.Done()
	for i := 0; i < quotaIterationsCount; i++ {
		fmt.Printf(formatWork(workerNum, i))

		/** Возарвщаем квоту обратно на каждой итерации квоты.
		Этот блок if позволит делится пропускной способностью(ресурсами программы) для всех воркеров и процедуры
		 будут проходить более плавно  */
		if i % quotaLimit == 0 {
			<-quotaCh
			quotaCh <- struct{}{}
		}
		runtime.Gosched()
 	}
	/** возвращаем слот. То есть мы читаем из этого канала и освобождаем там место для того, что бы другая горутина
		положила туда свой слот в этот канал */
	<-quotaCh


}
