package language

import (
	"fmt"
	"time"
)

func TimerStart(sec int) {
	t := time.Duration(sec)
	timer := time.NewTimer(time.Second * t)
	// multiplexer select
	select {
	case <- timer.C:
			// как только timer отсчитает t времени, отработает событие
			fmt.Println("timer.C, timeout happend")
	// короткое объявление new timer, когда нам возвращается канал без промежуточной переменной.
	// Но мы не можем остановить этот канал, так, что пока он не выполнится, мы не сможем освободить ресурсы,
	// даже если функция выполнится полностью
	case <-time.After(time.Minute):
		fmt.Println("time.After, timeout happend")
}
}
// Выгрузка ресурсов и отключение таймера
func timerStop(timer time.Timer) {
	if !timer.Stop() {
		<-timer.C
	}
}