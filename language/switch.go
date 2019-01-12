package language

import "fmt"

func TestSwitch() {
	printFormat := "| switch case | %v | \n"

	// 1. обычный switch case
	hashMap := make(map[string]string)
	hashMap["firstName"] = "Penny"
	hashMap["lastName"] = "Smite"
	switch hashMap["firstName"] {
	case "Tommy", "Ronny":
		fmt.Printf(printFormat, "Tommy or Ronny")
		// break не нужен он не провалится дальше
	case "Penny":
		fallthrough // эта конструкция используется, что бы провалиться дальше.
	default:
		fmt.Printf(printFormat, " default value")
	}

	// 2.  Альтернатива множественным if else
	switch {
	case hashMap["firstName"] == "Tommy":
		//
	case hashMap["lastName"] == "anybody":
		//
	}

	// 3. Если у нас вложенность Switch в цикл, тогда break остановит лишь switch
	// что бы этого избежать нужно использовать label
	MyLoop: // label
	for key, val := range hashMap {
		switch  {
		case key == "firstName" && val == "Penny" :
			// тормозим всю конструкцию
			break MyLoop
		default:
			fmt.Printf(printFormat, "Я никогда не выведусь")
		}
	}
}
