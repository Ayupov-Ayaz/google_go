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
	f.StartDefer()
	f.StartInit()
	f.StartPanicAndRecover()
	lang.RunFor()
	lang.TestAreaOfVisibility()	
	lang.TestPointers()
	types.TestStruct()
	duck := new(ent.Duck)
	lang.CookingDuck(duck)

	fmt.Println("\nПотоки: У Main всегда главный поток, который не ждет другие.")
	// Создание WaitGroup всегда по указателю, нет смысла его копировать
	wg := &sync.WaitGroup{}
	lang.TestGo(800, wg)
	lang.TestGo(300, wg)
	lang.TestGo(450, wg)
	lang.TestGo(100, wg)
	lang.TestGo(700, wg)
	wg.Wait()

}