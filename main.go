package main

import (
	"fmt"
	"net/http"
)

func PrintOnWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello :: This app is now deployed")
}

func main() {
	http.HandleFunc("/", PrintOnWeb)
	http.ListenAndServe(":8085", nil)
}
