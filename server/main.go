package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("r: ", r.Header)
	log.Println("resp: ", w.Header())
	w.WriteHeader(http.StatusOK)

}
