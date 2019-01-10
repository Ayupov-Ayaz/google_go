package language

import (
	"fmt"
	"google_go/Entities"
)

// приведение типов (расспознавание)
func StartTypeAssertions(i interface{}) {
	if s, ok := i.(string); ok {
		fmt.Printf("this is string: %s\n", s)
	}

	if d, ok := i.(int); ok {
		fmt.Printf("this is int %d \n", d)
	}

	if duck, ok := i.(Entities.Duck); ok {
		fmt.Printf("Это утка, утка умеет %s, %s, $s \n", duck)
		duck.Fly()
		duck.Quack()
		duck.Swim()
	}

}