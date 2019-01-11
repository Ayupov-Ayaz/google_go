package examples

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
 	Выводит на экран только уникальные значения введенных строк
 */
func Uniq(input io.Reader , writer io.Writer) error{
	// создаем сканер ввода
	in := bufio.NewScanner(input)
	var prev string
	// когда считывать будет нечего придет false
	for in.Scan() {
		txt := in.Text()
		if txt == prev { continue }
		if txt < prev {
			return fmt.Errorf("File not sorted!")
		}
		prev = txt
		fmt.Fprintln(writer, txt)
	}
	return nil
}

func StartUnique() {
	err := Uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}