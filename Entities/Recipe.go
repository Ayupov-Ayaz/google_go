package Entities
import "fmt"
type Recipe struct {
	Food  			   // Расширяем тип Recipe свойствами и методами типа Food
	Ingredients []Food // Агрегация  
}

func (this *Recipe) Eating() {
	for _, food := range this.Ingredients {
		fmt.Printf("(type Recipe) I eating %s \n", food.Name)
	}
}