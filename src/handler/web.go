package handler

import (
	"log"
	"net/http"

	"github.com/hasemeneh/golang-simple-login/helper/httphandler"
	"github.com/hasemeneh/golang-simple-login/helper/response"
	"github.com/hasemeneh/golang-simple-login/src/service"
	"github.com/julienschmidt/httprouter"
)

type WebService interface {
	Run()
}
type WebHandler struct {
	Service service.Service
	router  *httprouter.Router
}

func NewWebHandler(Service service.Service) WebService {
	return &WebHandler{
		router: httprouter.New(),
	}
}

func (w *WebHandler) Register() {
	apiHandlers := httphandler.New("/api", w.router)
	apiHandlers.GET("/ping", w.PING)
}

func (w *WebHandler) Run() {
	w.Register()
	log.Fatalln(http.ListenAndServe(":33339", w.router))
}
func (w *WebHandler) PING(r *http.Request) response.HttpResponse {
	return response.NewJsonResponse().SetMessage("Pong").SetData("Pung")
}
