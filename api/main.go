package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers func
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)

	/*
		handler => validation{1. request, 2. user} => business logic => response
		1. data model
		2. error handling
		3. session
	*/
}
