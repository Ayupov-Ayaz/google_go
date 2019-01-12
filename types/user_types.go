package types

import (
	"fmt"
	ent "google_go/Entities"
)

// Пользовательские типы
func TestUserTypes() {
	var ip = Ip{127,0,0,1} 
	fmt.Println(ip.correctIp())

	var word = Word("    Hello world!!    ")
	fmt.Println(*word.Trim().Wrap("|"))
}

type withFiles bool

func TestStruct() {
	apple := ent.Food{ "Apple"}
	tomato := ent.Food{"Tomato"}
	recipe := ent.Recipe{
		Ingredients: []ent.Food{apple, tomato},
	}
	recipe.Food.Name = "Pasta" // После расширения типа Recipe типом Food нам доступны все его методы и свойства

	// использование своих типов
	var withoutFiles = withFiles(true)
	fmt.Println(withoutFiles)

	// У Go lang отсутствует переопределение типов, так что можно вызывать методы
	// как своего типа так и типа на основе которого был расширен
	recipe.Food.Eating() // Запуск из типа Food
	recipe.Eating() // Запуск из типа Recipe
}
