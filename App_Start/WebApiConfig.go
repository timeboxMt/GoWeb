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

	fsh := http.FileServer(http.Dir("F:/goWork/GoWeb/template/Script"))
	http.Handle("/Script/", http.StripPrefix("/Script/", fsh))

	fshCss := http.FileServer(http.Dir("F:/goWork/GoWeb/template/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fshCss))

	http.HandleFunc("/login", controllers.Login)
}
