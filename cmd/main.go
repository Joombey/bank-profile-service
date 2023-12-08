package main

import (
	"farukh.go/profile/di"
	"farukh.go/profile/http"
)

func main() {
	di.Init()
	http.Run()
}
