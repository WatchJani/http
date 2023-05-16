package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Custom-Header", "Vrijednost 1")
		w.Header().Set("Custom-Header", "Vrijednost 2")

		w.Header().Del("Custom-Header")

		w.WriteHeader(http.StatusGatewayTimeout)
		w.Write([]byte("Janko"))

		cookies, err := r.Cookie("jsw")

		if err != nil {
			fmt.Println("MY BIG ERR", err)
		}

		fmt.Println(cookies.Value)

		fmt.Println(r.URL.Path)

	})

	http.ListenAndServe(":5000", nil)
}
