package main

import (
	"fmt"
	"net/http"

	"github.com/PineLover/golang_server_restapi/server/myapp"
)

func main() {
	fmt.Print("running")

	http.ListenAndServe(":3002", myapp.NewHttpHandler())
}
