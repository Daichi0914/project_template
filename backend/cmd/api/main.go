package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
