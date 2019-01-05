package types

import (
	"fmt"
)

// Пользовательские типы
func TestUserTypes() {
	var ip = Ip{127,0,0,1} 
	fmt.Println(ip.correctIp())

	var word = Word("    Hello world!!    ")
	fmt.Println(*word.Trim().Wrap("|"))
}
