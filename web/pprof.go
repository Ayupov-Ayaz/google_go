package web

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // импортируем, но не обращаемся к нему
	"time"
)
/**
 	Профилировщик Pprof используется для для поиска горячих мест в уже работающей программе под боевой нагрузкой
	использование:
	curl http://localhost:port/debug/pprof/heap - o mem_out.txt
	curl http://localhost:port/debug/pprof/profile?seconds=5 -o cpu_out.txt

	go tool pprof -sv -alloc_objects pprof_1.exe mem_out.txt > mem_ao.svg
	go tool pprof -svg pprof_1.exec cpu_out.txt > cpu.svg

 */
type Post struct {
	ID int
	Title string
	Text string
	Author string
	Comments int
	Time time.Time
}

 func PprofStart() {
 	http.HandleFunc("/", pprofTestHandle)
 	fmt.Println("starting serve :8000")
 	http.ListenAndServe(":8000", nil)
 }
 func pprofTestHandle(w http.ResponseWriter, r *http.Request) {
	s := ""
	for i := 0; i < 1000; i++ {
		p:= &Post{ID: i, Text: "new post"}
		s += fmt.Sprintf("%#v", p)
	}
	w.Write([]byte(s))
 }

