package main

import "google_go/web"

func main() {
	//
	go web.StartServeMux(8000)
	web.StartServeMux(8001)


}