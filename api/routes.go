package api

import (
	"net/http"
	"main/api/service"
)

func Routes(){
	http.HandleFunc("/",service.Get)

	http.HandleFunc("/post",service.Post)

	http.HandleFunc("/put/",service.Put)

	http.HandleFunc("/delete/",service.Delete)
}