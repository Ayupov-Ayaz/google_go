package language

import "fmt"

var (
	numb int8 = 12
)

const autoInt = 10
const staticInt int = 10

// iota - это автоинкремент для констант
const (
	one = iota + 1
	two
	_ // пропускаем значение
	four
)

const (
	_ = iota
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)


func TestConst() {
	fmt.Println("Работа с константами")
	var float float64 = 3.43
	fmt.Println( float + autoInt) // если у константы не указан тип, то он динамически будет подставляться
	fmt.Println(float + float64(staticInt)) // тут уже у константы тип указан
	fmt.Println(one, two, four)
	fmt.Printf("%d \n %d \n %d \n %d \n %d \n ", KB, MB, GB, TB, PB)
}