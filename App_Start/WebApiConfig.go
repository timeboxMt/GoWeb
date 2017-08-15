package main

import (
	"log"
	"net/http"

	"../controllers"
)

func main() {
	httpEntrance()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func httpEntrance() {
	//http.Handle("/static/", http.StripPrefix("/Script/", http.FileServer(http.Dir("Script"))))
	//http.Handle("/Script/", http.FileServer(http.Dir("Script")))

	fshjs := http.FileServer(http.Dir("F:/goWork/GoWeb/template/Script"))
	http.Handle("/Script/", http.StripPrefix("/Script/", fshjs))

	fshcss := http.FileServer(http.Dir("F:/goWork/GoWeb/template/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fshcss))

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/index", controllers.SayhelloName)
}
