package user

import (
	"context"

	"github.com/hasemeneh/golang-simple-login/src/models"
)

func (u *userUsecase) UpdateAddress(ctx context.Context, address, token string) (*models.UserModel, error) {
	user, err := u.GetUserInfoByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = u.User.UpdateAddress(ctx, address, user.ID)
	if err != nil {
		return nil, err
	}
	user.Address = address
	return user, nil
}

func (u *userUsecase) UpdatePassword(ctx context.Context, password, token string) (*models.UserModel, error) {
	user, err := u.GetUserInfoByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	err = u.User.UpdatePassword(ctx, password, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
