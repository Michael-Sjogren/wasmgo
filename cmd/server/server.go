package main

import (
	"fmt"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("../../assets")))

	if err != nil {
		fmt.Println("bruh u just crashed")
		return
	}

}
