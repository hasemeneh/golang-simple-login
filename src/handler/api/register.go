package api

import (
	"encoding/json"
	"net/http"

	"github.com/hasemeneh/golang-simple-login/helper/response"
	"github.com/hasemeneh/golang-simple-login/src/constants"
	"github.com/hasemeneh/golang-simple-login/src/models"
)

func (p *Public) Register(r *http.Request) response.HttpResponse {
	var request models.UserModel
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusBadRequest))
	}

	err = p.Service.UserUsecase.Register(r.Context(), &request)
	if err != nil {
		if err == constants.ErrEmailAlreadyRegistered || err == constants.ErrUsernameAlreadyUsed {
			return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusBadRequest))
		}
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusInternalServerError))
	}
	return response.NewJsonResponse().SetMessage("Registered")
}
