package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Login 登陆方法
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../View/login.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		//跳转到主页
		t, _ := template.ParseFiles("../View/index.html")
		log.Println(t.Execute(w, nil))
	}
}
