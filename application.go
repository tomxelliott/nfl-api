package main

import (
	"net/http"

	"log"
	"runtime"
	"os"
	"fmt"

	"portfolio/logic"
)

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	port := os.Getenv("PORT")
	fmt.Print(port)
	if port == "" {
		port = "5000"
	}

	r := logic.NewRouter()
	http.ListenAndServe(":" + port, r)
}
