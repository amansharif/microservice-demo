package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func formFillUp(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("postrequest.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}


func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/formFillup", formFillUp)
	http.ListenAndServe(":80",nil)

}