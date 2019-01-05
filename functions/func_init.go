package functions

import (
	"fmt"
	"google_go/subpackage"
)

func init() {
	fmt.Println("call function init from func_init file")
}

func StartInit() {
	fmt.Println("function StartInit from func_init file")
	subpackage.Use()
}