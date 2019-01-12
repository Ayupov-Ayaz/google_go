package types

import "fmt"

func TestSlice() {
	var printFormatDelimetr string= "| slice | %v | %s | \n"
	// Срезы

	// 1 способ, создание среза с указанием емкости среза(15) и текущей длины (5)
	first := make([]byte, 5, 15)
	fmt.Printf("first slice len(%d) cap(%d) = %v \n", len(first), cap(first), first)

	// 2 способ, создание среза с указанием текущей длины, емкость такая же как и длина
	second := make([]uint8, 3)
	fmt.Printf("second slice len(%d) cap(%d) = %v \n", len(second), cap(second), second)

	// 3 способ, создание среза непосредственно из массива.
	var arr = [4]byte{1, 2, 3, 4}
	// [1:4]
	third := arr[1:4]
	//third := arr[:3]
	//third := arr[:]
	fmt.Println(third)

	// значения в срез передаются по ссылке! при изменеиии в срезе, поменяются и в массиве
	third[0] = 123
	fmt.Println(third)
	fmt.Println(arr)

	// Копирование срезов
	// Для копирования одного среза в другой обязательно нужно указать длину при создании среза куда будет
	// всё скопировано
	CopyThird := make([]byte, len(third))
	fmt.Printf(printFormatDelimetr, copy(CopyThird, third), "Элемента скопировано")
	// Копирование таким способом позволит нам передать элементы по значению, а не по ссылке
	third = []byte{2,3,1,6,2}
	fmt.Println("1) len-" , len(third), "cap-", cap(third)) // len- 5, cap- 5

	fmt.Println(CopyThird)
	fmt.Println(third)

	// Расширение срезов
	// 1. Добавление одного среза в другой. Троеточие в конце обязательно!
	third = append(third, CopyThird[:]...)
	fmt.Println("2) len-" , len(third), "cap-", cap(third)) // len - 12, cap - 16
	fmt.Println(third)

	// 2. Перечисление доп элементов через запятую
	third = append(third, 100, 101, 102, 103)
	fmt.Println("3) len-" , len(third), "cap-", cap(third)) // len- 12, cap- 16
	fmt.Println(third)
	// Емкость задалось с запасом

}