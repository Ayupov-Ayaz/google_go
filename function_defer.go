package main

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
func first() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("fitst")
}

func second() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("second")
}

func third() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("third")
}

func fourth() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("fourth")
}

func fifth() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("fifth")
}

func main() {
	fmt.Println("begin main function")
	defer first()
	defer second()
	defer third()
	defer fourth()
	defer fifth()
	fmt.Println("end main function")

}