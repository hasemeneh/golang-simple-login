package repositories

import (
	"context"

	"github.com/hasemeneh/golang-simple-login/src/models"
)

type UserDoman interface {
	InsertUser(ctx context.Context, userModel *models.UserModel) error
	UpdateAddress(ctx context.Context, address string, UserID int64) error
	UpdatePassword(ctx context.Context, password string, UserID int64) error
	GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error)
	GetUserByEmail(ctx context.Context, email string) (*models.UserModel, error)
	GetUserByID(ctx context.Context, id int64) (*models.UserModel, error)
}
