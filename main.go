package main

import (
	"fmt"
	ent "google_go/Entities"
	f "google_go/functions"
	lang "google_go/language"
	"google_go/types"
	"sync"
)

func main() {
	lang.TestSwitch()
	f.StartDefer()
	f.StartInit()
	f.StartPanicAndRecover()
	lang.TestRune()
	lang.RunFor()
	types.TestMaps()
	lang.TestAreaOfVisibility()
	lang.TestPointers()
	types.TestStruct()
	duck := new(ent.Duck)
	lang.CookingDuck(duck)
	lang.TestConst()
	types.TestSlice()
	fmt.Println("\nПотоки: У Main всегда главный поток, который не ждет другие.")
	// Создание WaitGroup всегда по указателю, нет смысла его копировать
	wg := &sync.WaitGroup{}
	lang.TestGo2(3, wg)
	lang.TestChannel()
	wg.Wait()
}