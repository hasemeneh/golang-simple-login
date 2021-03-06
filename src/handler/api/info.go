package api

import (
	"net/http"
	"strings"

	"github.com/hasemeneh/golang-simple-login/helper/response"
	"github.com/hasemeneh/golang-simple-login/src/constants"
)

const authPrefix string = "Bearer "

func (p *Public) HandleGetUserInfo(r *http.Request) response.HttpResponse {
	var token string
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		token = strings.Replace(authHeader, authPrefix, "", 1)
	}

	resp, err := p.Service.UserUsecase.GetUserInfoByToken(r.Context(), token)
	if err != nil {
		if err == constants.ErrInvalidToken {
			return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusUnauthorized))
		}
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusInternalServerError))
	}
	return response.NewJsonResponse().SetData(resp)
}
