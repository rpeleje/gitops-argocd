package main

import "net/http"

func main() {
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("<h1>Hello Peleje!!!</h1> <h2>Sync manual blablabl!</h2>"))
	})
	http.ListenAndServe(":4001", nil)
}