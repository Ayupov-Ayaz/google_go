package main

import(
	f "google_go/functions"
	lang "google_go/language"
	types "google_go/types"
	ent "google_go/Entities"
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