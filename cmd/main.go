package main

import (
	// "os"

	// cts "farukh.go/profile/constants"
	// "farukh.go/profile/dao/db"
	"farukh.go/profile/http"
	// in "farukh.go/profile/internal"
)

func main() {
	// if path := os.Getenv("CONFIG_PATH"); path == "" {
	// 	in.Init(cts.LocalConfigPath)
	// } else {
	// 	in.Init(path)
	// }
	// db.Init()
	http.RunYan()
}
