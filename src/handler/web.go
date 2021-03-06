package handler

import (
	"log"
	"net/http"

	"github.com/hasemeneh/golang-simple-login/helper/httphandler"
	"github.com/hasemeneh/golang-simple-login/src/handler/api"
	"github.com/hasemeneh/golang-simple-login/src/service"
	"github.com/julienschmidt/httprouter"
)

type WebService interface {
	Run()
}
type WebHandler struct {
	Service *service.Service
	router  *httprouter.Router
}

func NewWebHandler(Service *service.Service) WebService {
	return &WebHandler{
		router:  httprouter.New(),
		Service: Service,
	}
}

func (w *WebHandler) Register() {
	apiHandlers := api.GetHandler(w.Service)
	apiHttpHandlers := httphandler.New("/api", w.router)
	apiHttpHandlers.GET("/ping", apiHandlers.PING)
	apiHttpHandlers.GET("/user", apiHandlers.HandleGetUserInfo)
	apiHttpHandlers.POST("/register", apiHandlers.Register)
	apiHttpHandlers.POST("/login", apiHandlers.Login)
	apiHttpHandlers.POST("/address", apiHandlers.HandleUpdateAddress)
	apiHttpHandlers.POST("/change_password", apiHandlers.HandleUpdatePassword)
}

func (w *WebHandler) Run() {
	w.Register()
	log.Println("Server running on localhost:33339")
	log.Fatalln(http.ListenAndServe(":33339", w.router))
}
