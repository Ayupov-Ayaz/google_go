package main

import (
	"fmt"
	"google_go/subpackage"
)

func init() {
	fmt.Println("call function init from function_init file")
}

func main() {
	fmt.Println("function main from function_init file")
	subpackage.Use()
}