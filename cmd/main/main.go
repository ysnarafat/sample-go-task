package main

import (
	"fmt"
	"html/template"
	"net/http"
	"xyz/test"
)

func main() {

	test.GetData()

	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}
