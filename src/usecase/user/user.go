package user

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/hasemeneh/golang-simple-login/helper/redis"
	"github.com/hasemeneh/golang-simple-login/src/constants"
	"github.com/hasemeneh/golang-simple-login/src/models"
	"github.com/hasemeneh/golang-simple-login/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	User  repositories.UserDoman
	Redis redis.RedisInterface
}

func New(user repositories.UserDoman, Redis redis.RedisInterface) *userUsecase {

	return &userUsecase{
		User:  user,
		Redis: Redis,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *models.UserModel) error {
	_, err := u.User.GetUserByEmail(ctx, user.Email)
	if err != sql.ErrNoRows {
		if err == nil {
			return constants.ErrEmailAlreadyRegistered
		}
		return err
	}
	_, err = u.User.GetUserByUsername(ctx, user.Username)
	if err != sql.ErrNoRows {
		if err == nil {
			return constants.ErrUsernameAlreadyUsed
		}
		return err
	}
	return u.User.InsertUser(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.User.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	isValid, err := u.ValidatePassword(ctx, user, password)
	if err != nil {
		return "", err
	}
	if !isValid {
		return "", constants.ErrWrongPassword
	}
	token := generateToken(user.ID)
	_, err = u.Redis.Set(fmt.Sprintf(constants.SessionToken, token), fmt.Sprintf("%d", user.ID), time.Hour*24)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *userUsecase) ValidatePassword(ctx context.Context, user *models.UserModel, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func generateToken(userID int64) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", userID)))
	h.Write([]byte(time.Now().Format(time.RFC3339Nano)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

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
