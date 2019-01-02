package main

import "fmt"

// Неупорядоченные хэш таблицы

func main() {
	// Способы объявить карту:

	// 1. С указанием длины карты
	first := make(map[bool]int8, 2) // ключи имеют тип bool, значения int8

	// 2. Без указания длины карты
	strBool := make(map[string]bool)

	// 3. Инициализация + присвоение данных
	ages := map[string]uint{"Tommy" : 27, "Ron" : 20, "Andy" : 23}

	// 4. Можно создать пустую карту
	emptyMap := map[string]uint{}

	first[false] = 34
	first[true] = 55
	strBool["empty"] = true
	strBool["came"] = false
	emptyMap["one"] = 1
	emptyMap["two"] = 2

	// 5. Удаление
	delete(emptyMap,"one")

	fmt.Println(first)
	fmt.Println(strBool)
	fmt.Println(ages)
	fmt.Println(emptyMap)
}