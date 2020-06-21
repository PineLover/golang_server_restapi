package main

import (
	"net/http"

	"github.com/PineLover/golang_server_restapi/server/myapp"
)

func main(){
	http.ListenAndServe(":3000", myapp.NewHandler())
}