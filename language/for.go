package language

import (
	"fmt"
)

func RunFor() {
	var arr = []string {
		"Яблоко", "Груша", "Ананас", "Помидор", "Картошка", "Хлеб",
	}

	var menu = make(map[string]string)
	menu["fruit"] = "Apple"
	menu["meet"] = "Beef"
	menu["drink"] = "Green Tea"
	menu["first"] = "noodles"
	menu["second"] = "rise"

	// 1. 
	fmt.Println("for 1.")
	for i := 0; i < len(arr); i++ {
		//
	}

	// 2.
	i := 0
	for i < len(arr) {
		//
		i++
	}

	// 3. можно также запустить без каких либо аргументов цикла
	i = 0
	for {
		if(i < len(arr)) {
			break
		}
		i++

	}
	// 4. аналог foreach, очень полезно использовать с Map
	for key, val := range menu {
		fmt.Println(key , val)
	}

	// 5. можно получить лишь значения без ключей
	for _, val := range menu {
		fmt.Println(val)
	}
}
