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

	http.HandleFunc("/login", controllers.Login)
}
