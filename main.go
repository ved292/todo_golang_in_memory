package main

import (
	"log"
	"net/http"
	"main/routes"
)

func main(){
	http.HandleFunc("/",routes.Get)

	http.HandleFunc("/post",routes.Post)

	http.HandleFunc("/put/",routes.Put)

	http.HandleFunc("/delete/",routes.Delete)
	
	log.Fatal(http.ListenAndServe(":8000",nil))
}