package language

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalOperations2 int32
func StartRaceConditionRelationWithAtomic() {
		for i := 0; i < 1000; i++ {
			go inc2()
		}
		time.Sleep(time.Millisecond * 2)
		fmt.Println(totalOperations2)
}

func inc2() {
	atomic.AddInt32(&totalOperations2, 1) //Атомарное увеличение счетчика
}