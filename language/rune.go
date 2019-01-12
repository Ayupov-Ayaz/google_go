package language

import "fmt"

// unicode символы
func TestRune() {
	printFormat := "| Runes | %v | %v | %v | \n"
	var symbol rune = 'a'
	var unicodeRune rune = 'Ʊ'
	var unicodeSymbolByNumber rune = '\u0026'

	fmt.Printf(printFormat, symbol, unicodeRune, unicodeSymbolByNumber)
}