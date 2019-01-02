package main

import "fmt"

func main() {
	first := "Hello, \n World!!!!"
	second := `Goodbye, \n World!!!`
	fmt.Println(first)
	fmt.Println(second)

	//1. При обращении к строке как к массиву мы получим unicode code символа
	// что бы преобразовать код в строку необходимо кастануть его
	fmt.Println(string(first[0]))

	//2. можно обращаться к строкам как к срезам
	fmt.Println(first[0:3])

	//3. Работа с многобайтными строками
	rus := "А тут возникнут сложности. Необходимо преобразовывать строки к rune"
	// это позволит обращаться к каждому символу отдельно, а не к байту
	runeRus := []rune(rus)
	fmt.Println(string(runeRus[27:]))





}