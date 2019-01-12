package types

import "fmt"

// Создание своих типов ошибок
type ErrorString string

func (e ErrorString) Error() string{
	return string(e)
}

// Своя структура с подробным описанием ошибки
type MyError struct {
	Msg string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d %s", e.File, e.Line, e.Msg)
}

func BadFunc() error {
	return &MyError{Msg:"Какая-то ошибка", File:"error.go", Line:24,}
}
