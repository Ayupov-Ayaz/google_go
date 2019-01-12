package language

import (
	"fmt"
	"google_go/types"
)

func StartErrors() {
	err := types.BadFunc()
	// проверка на type может существовать только в блоке switch
	switch err := err.(type) {
	case nil:
		// всё ок, ошибки нет
	case *types.MyError:
		fmt.Println(err.Line, err.File, err.Msg)
	default:
		// unknown error
	}

}