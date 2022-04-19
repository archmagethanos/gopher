package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/archmagethanos/gopher/handlers"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.ListenAndServe(":9090", nil)
}
