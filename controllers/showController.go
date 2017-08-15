package controllers

import (
	"html/template"
	"net/http"
)

//SayhelloName 输出Hello
func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../View/index.html")
		t.Execute(w, nil)
	}
}
