package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Custom-Header", "Vrijednost 1")
		// w.Header().Set("Custom-Header", "Vrijednost 2")

		w.Header().Set("X-Forwarded-For", "12121.1.321.21.21.321")

		// w.Header().Del("Custom-Header")

		// w.WriteHeader(http.StatusGatewayTimeout)

		// w.Write([]byte("Janko"))
		err := r.ParseForm()
		if err != nil {
			// Greška pri parsiranju forme
			http.Error(w, "Greška pri parsiranju forme", http.StatusBadRequest)
			return
		}

		// cookies, err := r.Cookie("jsw")

		// if err != nil {
		// 	fmt.Println("MY BIG ERR", err)
		// }

		// fmt.Println(cookies.Value)

		fmt.Println(r.URL.Path)

		// get, _, err := r.FormFile("name")

		// if err != nil {
		// 	fmt.Println(err)
		// }
		// name := r.Form.Get("name")

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Greška pri čitanju tijela zahtjeva", http.StatusBadRequest)
		// 	return
		// }

		fmt.Println(r.Host)

		// Ispisivanje tijela zahtjeva
		// fmt.Println("Tijelo zahtjeva:", string(body))

		fmt.Println(r.ContentLength)
		// fmt.Println(name)

		fmt.Println(r.RemoteAddr)

	})

	http.ListenAndServe(":5000", nil)
}
