package main

import "fmt"

func main(){
	// несколько способов объявить массивы
	// 1.
	var arr [3]int8
	arr[1] = 23
	// 2.
	arr2 := [5]uint{1,2,3,4}
	arr2[1] = 245
	// 3. сам определелит сколько элетентов в массиве и создаст нужную длину
	arr3 := [...]int16{123,43,554}
	arr3[2] = 47474

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

}
