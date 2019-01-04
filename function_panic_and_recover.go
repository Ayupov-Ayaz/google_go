package main

import (
	"fmt"
)
var startFunc string = "Starting function \"%s\"  \n"
var recoverFuncStr string = "recovery exception called from  function %s : %s \n"
var issetError = "Я не выполнюсь, если найдется ошибка"
func first() {
	fmt.Printf(startFunc, "first")
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
	fmt.Printf(startFunc, "second")
		
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
	
	fmt.Printf(startFunc, "third")
	
	// recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf(recoverFuncStr, "third", error)
			panic(error)
		}
	}()	
	
	panic("someone Exception")
}

func main() {
	
	fmt.Printf(startFunc, "main")
	// recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf(recoverFuncStr, "Main", error)
		}
	}()
	first()
	// Этот код не выполнится, так как обработчик ошибок остановит программу
	fmt.Println(issetError)
}
