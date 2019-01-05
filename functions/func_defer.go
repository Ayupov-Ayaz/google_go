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
func _first() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("fitst")
}

func _second() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("_second")
}

func _third() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("_third")
}

func _fourth() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("_fourth")
}

func _fifth() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("_fifth")
}

func StartDefer() {
	fmt.Println("begin start function")
	defer _first()
	defer _second()
	defer _third()
	defer _fourth()
	defer _fifth()
	fmt.Println("end start function")

}