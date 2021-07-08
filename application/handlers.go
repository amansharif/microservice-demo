package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, formFillUp)
}

func formFillUp(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("postrequest.html")
	if err != nil {
		fmt.Println(err)
	}
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	details := data{
		Firstname: r.FormValue("fname"),
		Lastname:  r.FormValue("lname"),
		Age:       r.FormValue("age"),
		Gender:    r.Form.Get("gender"),
		Comment:   r.FormValue("comment"),
	}

	saveData(details)

	t.Execute(w, struct{ Success bool }{true})
}

func getResponse(w http.ResponseWriter, r *http.Request) {
	myValues := getData()
	fmt.Println(myValues)
	t, err := template.ParseFiles("getresponse.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, myValues)
}
//func getRequest(w http.ResponseWriter, r *http.Request) {
//	t, err := template.ParseFiles("getrequest.html")
//	if err != nil {
//		fmt.Println(err)
//	}
//	t.Execute(w, nil)
//}
