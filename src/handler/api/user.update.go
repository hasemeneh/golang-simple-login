package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hasemeneh/golang-simple-login/helper/response"
	"github.com/hasemeneh/golang-simple-login/src/models"
)

func (p *Public) HandleUpdateAddress(r *http.Request) response.HttpResponse {
	var token string
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		token = strings.Replace(authHeader, authPrefix, "", 1)
	}
	var request models.UserModel
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusBadRequest))
	}

	resp, err := p.Service.UserUsecase.UpdateAddress(r.Context(), request.Address, token)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusInternalServerError))
	}
	return response.NewJsonResponse().SetData(resp)
}

func (p *Public) HandleUpdatePassword(r *http.Request) response.HttpResponse {
	var token string
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		token = strings.Replace(authHeader, authPrefix, "", 1)
	}
	var request models.UserModel
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusBadRequest))
	}

	resp, err := p.Service.UserUsecase.UpdatePassword(r.Context(), request.PlainPassword, token)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusInternalServerError))
	}
	return response.NewJsonResponse().SetData(resp)
}
