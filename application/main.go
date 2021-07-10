package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/formFillup", formFillUp)
	http.HandleFunc("/details", getResponse)
	//http.HandleFunc("/entername", getRequest)
	log.Fatal(http.ListenAndServe(":80", nil))
	//lets learn git
}
