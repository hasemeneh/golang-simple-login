package user

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hasemeneh/golang-simple-login/src/constants"
	"github.com/hasemeneh/golang-simple-login/src/models"
)

func (u *userUsecase) GetUserInfoByToken(ctx context.Context, token string) (*models.UserModel, error) {
	isExist, err := u.Redis.Exist(fmt.Sprintf(constants.SessionToken, token))
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, constants.ErrInvalidToken
	}
	UserIDString, err := u.Redis.Get(fmt.Sprintf(constants.SessionToken, token))
	if err != nil {
		return nil, err
	}
	userID, err := strconv.ParseInt(UserIDString, 10, 64)
	if err != nil {
		return nil, err
	}
	userData, err := u.User.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
