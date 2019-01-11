package main

import (
	"google_go/Entities"
	"google_go/functions"
	lang "google_go/language"
	"google_go/types"
)

func main() {
	lang.TestSwitch()
	functions.StartDefer()
	functions.StartInit()
	functions.StartPanicAndRecover()
	lang.TestRune()
	lang.RunFor()
	types.TestMaps()
	lang.TestAreaOfVisibility()
	lang.TestPointers()
	types.TestStruct()
	duck := new(Entities.Duck)
	lang.CookingDuck(duck)
	lang.TestConst()
	types.TestSlice()
	lang.TestChannel()
	lang.RunGoroutines()
}