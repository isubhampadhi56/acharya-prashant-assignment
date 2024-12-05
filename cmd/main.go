package main

import (
	"fmt"
	"net/http"

	router "github.com/api-assignment/pkg/routes"
)

func main() {
	fmt.Println("Hello World")
	router := router.MainRouter()
	http.ListenAndServe(":3000", router)
}
