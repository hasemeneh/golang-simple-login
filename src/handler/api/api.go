package api

import (
	"net/http"

	"github.com/hasemeneh/golang-simple-login/helper/response"
	"github.com/hasemeneh/golang-simple-login/src/service"
)

type Public struct {
	Service *service.Service
}

func GetHandler(Service *service.Service) *Public {
	return &Public{
		Service: Service,
	}
}

func (w *Public) PING(r *http.Request) response.HttpResponse {
	return response.NewJsonResponse().SetMessage("Pong").SetData("Pung")
}
