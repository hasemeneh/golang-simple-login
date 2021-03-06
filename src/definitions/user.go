package definitions

import (
	"context"

	"github.com/hasemeneh/golang-simple-login/src/models"
)

type UserUsecase interface {
	Register(ctx context.Context, user *models.UserModel) error
	Login(ctx context.Context, username, password string) (string, error)
	GetUserInfoByToken(ctx context.Context, token string) (*models.UserModel, error)
}
