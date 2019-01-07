package main

import (
	ent "google_go/Entities"
	f "google_go/functions"
	lang "google_go/language"
	"google_go/types"
)

func main() {
	f.StartDefer()
	f.StartInit()
	f.StartPanicAndRecover()
	lang.RunFor()
	lang.TestAreaOfVisibility()	
	lang.TestPointers()
	types.TestStruct()
	duck := new(ent.Duck)
	lang.CookingDuck(duck)
}