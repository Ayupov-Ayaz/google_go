package types

import "fmt"

// Неупорядоченные хэш таблицы

func TestMaps() {
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
	strMap := map[string]string{}
	strMap["Hello"] = "Привет"

	// 5. Удаление
	delete(emptyMap,"one")

	// 6. Если обратиться к ключу которого не существует, то map нам выдаст значение по умолчанию типа по которому создана map
	fmt.Println(ages["Ayaz"], strMap["Hey"], strBool["isset"] ) // пустую строку не видно, но она там есть)
	// что бы понять был ли этот ключ установлен в мапе или нет можно воспользоваться вторым переменным

	// проверка на существование значения в map
	if _ , ok := strMap["Hey"]; ok {
		fmt.Println("Значение установлено")
	} else {
		fmt.Println("Вы получили значение по умолчанию, по ключу Hey в map strMap значения нету")
	}

	// 7. Копирование маp-ов только через цикл
	fmt.Println(first)
	fmt.Println(strBool)
	fmt.Println(ages)
	fmt.Println(emptyMap)
}