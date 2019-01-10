package language

import (
	"fmt"
	"google_go/Interfaces"
)
// в Go нигде не указывается, что тип имплементирует интерфейс
// Тут идет принцип Утиной типизации
func CookingDuck(duck Interfaces.DuckFamily) {
	fmt.Println("Если наш тип умеет:")
	duck.Swim()
	duck.Fly()
	duck.Quack()
	fmt.Println("Тогда это утка! Весь принцип интерфесов в языке Go-lang построен на утиной типизации")
	GoFly(duck)
}

// Проверка на то, что наш тип интерпритирует интерфейс
func GoFly(i interface{}) {
	if flyer, ok := i.(Interfaces.Flyer); ok {
		flyer.Fly()
	}
}
