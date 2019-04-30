package main

import (
	"go-sample/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", utils.Logging(HomeHandle, "home"))
	router.GET("/todos", utils.CommonHeaders(TodoHome, "todo-home"))
	http.ListenAndServe(":8080", router)
}
