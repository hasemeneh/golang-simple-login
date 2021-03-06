package main

import (
	"github.com/hasemeneh/golang-simple-login/src/handler"
	"github.com/hasemeneh/golang-simple-login/src/service"
)

func main() {
	serviceObj := service.Service{}
	web := handler.NewWebHandler(serviceObj)
	web.Run()
}
