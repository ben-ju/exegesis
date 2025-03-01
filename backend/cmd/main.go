package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world !")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil))
}
