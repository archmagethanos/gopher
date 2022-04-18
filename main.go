package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello World")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Oops"))
		}
		log.Printf("Data %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye")
	})

	http.ListenAndServe(":9090", nil)
}
