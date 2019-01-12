package functions

import (
	"fmt"
)
/**
 	Использование функции panic очень затратно, сильно подтормаживает программу
 */
var startFunc string = "Starting function \"%s\"  \n"
var recoverFuncStr string = "recovery exception called from  function %s : %s \n"
var issetError = "Я не выполнюсь, если найдется ошибка"

func first() {
		// recover
		defer func() {
			if error := recover(); error != nil {
				fmt.Printf(recoverFuncStr, "first", error)
				panic(error)
			}
		}()

	second()
	fmt.Println(issetError)

}

func second() {
	// recover
		defer func() {
			if error := recover(); error != nil {
				fmt.Printf(recoverFuncStr, "second", error)
				panic(error)
			}
		}()
	third()
	fmt.Println(issetError)
	
}

func third() {
	// recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf(recoverFuncStr, "third", error)
			panic(error)
		}
	}()	
	
	panic("someone Exception")
}

func StartPanicAndRecover() {
	// recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf(recoverFuncStr, "StartPanicAndRecover", error)
		}
	}()
	first()
	// Этот код не выполнится, так как обработчик ошибок остановит программу
	fmt.Println(issetError)
}
