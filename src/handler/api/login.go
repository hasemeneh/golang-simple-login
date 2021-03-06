package api

import (
	"encoding/json"
	"net/http"

	"github.com/hasemeneh/golang-simple-login/helper/response"
	"github.com/hasemeneh/golang-simple-login/src/constants"
	"github.com/hasemeneh/golang-simple-login/src/models"
)

func (p *Public) Login(r *http.Request) response.HttpResponse {
	var request models.UserModel
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusBadRequest))
	}

	token, err := p.Service.UserUsecase.Login(r.Context(), request.Username, request.PlainPassword)
	if err != nil {
		if err == constants.ErrWrongPassword {
			return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusUnauthorized))
		}
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusInternalServerError))
	}
	return response.NewJsonResponse().SetData(models.LoginResponse{Token: token})
}
