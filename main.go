package main

import (
	"log"
	"net/http"
	"main/api"
)

func main(){
	api.Routes()
	
	log.Fatal(http.ListenAndServe(":8000",nil))
}