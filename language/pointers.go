package language

import(
	"fmt"
	"strconv"
)

/**
	Указатели - это переменные которые хранят ссылку в область памяти
*/
func TestPointers() {
	var titleSpaceCount = 14
	var tableSpaceCount int = 15
	var space string
	
	var i int = 0
	for i < 16 {
		space += "-"
		i++ 
	}
	
	printSpace := func() {	
		fmt.Printf("|%s |%s|%s|%s|\n",space,space,space,space)
	}

	getTableTitleFormat := func() string {
		count := strconv.Itoa(titleSpaceCount)
		return "| %-" + count + "s | %-14s | %-14s | %-14s |\n"
	}

	getTableFormat := func() string {
		count := strconv.Itoa(tableSpaceCount)
		return "| %-" + count + "v | %-14T | %-14v | %-14p |\n";
	}

	fmt.Printf(getTableTitleFormat(), "Имя переменной ", "Тип переменной", "Значение", "Адрес")
	printSpace()	
	var intVar int8 = 10
	fmt.Printf(getTableFormat(), "intVar ", intVar, intVar, &intVar )
	printSpace()

	firstPointer := &intVar
	fmt.Printf(getTableFormat(), "firstPointer", firstPointer, firstPointer, &firstPointer)
	fmt.Printf(getTableFormat(), "*firstPointer", *firstPointer, *firstPointer, &*firstPointer)
	printSpace()

	secondPointer := &firstPointer
	fmt.Printf(getTableFormat(), "secondPointer", secondPointer, secondPointer, &secondPointer)
	fmt.Printf(getTableFormat(), "*secondPointer", *secondPointer, *secondPointer, &*secondPointer)
	fmt.Printf(getTableFormat(), "**secondPointer", **secondPointer, **secondPointer, &**secondPointer)
	printSpace()

	
	var thirdPointer *int
	fmt.Printf(getTableFormat(), "thirdPointer", thirdPointer, thirdPointer, &thirdPointer)
	printSpace()

}
