package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "＼はいさい、Hello World／")
}

func main() {
	log.Print("test print")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
