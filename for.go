package main
import(
	"fmt"
)

func main() {
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
		fmt.Println(arr[i])
	}
	
	fmt.Println("\nfor 2.")
	i := 0
	for i < len(arr) {
		fmt.Println(arr[i])
		i++
	}

	// можно также запустить без каких либо аргументов цикла
	fmt.Println("\nfor 3.")
	i = 0
	for {
		if(i < len(arr)) {
			break
		}
		fmt.Println(arr[i])
		i++

	}
	// аналог foreach, очень полезно использовать с Map
	fmt.Println("\n for range 1")
	for key, val := range menu {
		fmt.Println(key , val)
	}

	// можно получить лишь значения без ключей
	fmt.Println("\n for range 2")
	for _, val := range menu {
		fmt.Println(val)
	}
	fmt.Println(concatString("Привет", " ", "от", " функции", " concatString"))
}

func concatString(args ... string) string {
	var result string = ""
	for _, val := range args {
		result += val
	}
	return result
}