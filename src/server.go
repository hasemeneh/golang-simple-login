package main

import (
	"github.com/hasemeneh/golang-simple-login/src/handler"
	"github.com/hasemeneh/golang-simple-login/src/service"
)

func main() {
	web := handler.NewWebHandler(service.New())
	web.Run()

}
