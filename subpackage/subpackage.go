package subpackage

import "fmt"

func init() {
	fmt.Println("call function init from subpackage.go file")
}

// public function, так как начинается с заглавной буквы
func Use() {
	fmt.Println("Used subpackage")
}