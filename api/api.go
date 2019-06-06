package main

import (
	"net/http"
	"api/routers"
)

func main() {
	http.ListenAndServe(":5000", routers.GetRouter())
}