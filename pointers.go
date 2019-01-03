package main

import(
	"fmt"
	"strconv"
)

/**
	Указатели - это переменные которые хранят ссылку в области памяти
*/
func main() {
	var spaceCount int8 = 16
	
	fmt.Printf(getTableTitleFormat(14), "Имя переменной ", "Тип переменной", "Значение", "Адрес")
	printSpace(spaceCount)	
	var intVar int8 = 10
	fmt.Printf(getTableFormat(15), "intVar ", intVar, intVar, &intVar )
	printSpace(spaceCount)

	firstPointer := &intVar
	fmt.Printf(getTableFormat(15), "firstPointer", firstPointer, firstPointer, &firstPointer)
	fmt.Printf(getTableFormat(15), "*firstPointer", *firstPointer, *firstPointer, &*firstPointer)
	printSpace(spaceCount)

	secondPointer := &firstPointer
	fmt.Printf(getTableFormat(15), "secondPointer", secondPointer, secondPointer, &secondPointer)
	fmt.Printf(getTableFormat(15), "*secondPointer", *secondPointer, *secondPointer, &*secondPointer)
	fmt.Printf(getTableFormat(15), "**secondPointer", **secondPointer, **secondPointer, &**secondPointer)
	printSpace(spaceCount)

	
	var thirdPointer *int
	fmt.Printf(getTableFormat(15), "thirdPointer", thirdPointer, thirdPointer, &thirdPointer)
	printSpace(spaceCount)
}

func printSpace(spaceCount int8) {
	var space string
	var i int8 = 0
	for i < spaceCount{
		space += "-"
		i++ 
	}

	fmt.Printf("|%s |%s|%s|%s|\n",space,space,space,space)
}

func getTableFormat(spaceCount int) string {
	count := strconv.Itoa(spaceCount)
	return "| %-" + count + "v | %-14T | %-14v | %-14p |\n";
}

func getTableTitleFormat(spaceCount int) string {
	count := strconv.Itoa(spaceCount)
	return "| %-" + count + "s | %-14s | %-14s | %-14s |\n"
}