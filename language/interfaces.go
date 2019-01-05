package language
import(
	"fmt"
	interfaces "google_go/Interfaces"
)
// в Go нигде не указывается, что тип имплементирует интерфейс
// Тут идет принцип Утиной типизации
func CookingDuck(duck interfaces.DuckFamily) {
	fmt.Println("Если наш тип умеет:")
	duck.Swim()
	duck.Fly()
	duck.Quack()
	fmt.Println("Тогда это утка! Весь принцип интерфесов в языке Go-lang построен на утиной типизации")
}