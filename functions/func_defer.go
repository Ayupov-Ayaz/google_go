package functions

import (
	"fmt"
	"time"
)
/**
	Функция defer это отложенный запуск
	Используется для обязательного завершения каких нибудь процессов, 
	например: закрытия подключения к бд, закрытия файла для чтения
	работает по принципе stack (LIFO)
*/
var printFormat = "| func_defer | %v | \n"
func _first() {
	//time.Sleep(100 * time.Millisecond)
	fmt.Printf(printFormat,"fitst")
}

func _second() {
	//time.Sleep(500 * time.Millisecond)
	fmt.Printf(printFormat,"_second")
}

func _third() {
	//time.Sleep(1000 * time.Millisecond)
	fmt.Printf(printFormat,"_third")
}

func _fourth() {
	//time.Sleep(500 * time.Millisecond)
	fmt.Printf(printFormat,"_fourth")
}

func _fifth() {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf(printFormat,"_fifth")
}

func StartDefer() {
	fmt.Printf(printFormat,"begin StartDefer function")
	defer _first()
	defer _second()
	defer _third()
	defer _fourth()
	defer _fifth()
	fmt.Printf(printFormat,"end StartDefer function")

}