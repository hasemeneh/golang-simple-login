package user

import (
	"github.com/hasemeneh/golang-simple-login/helper/redis"
	"github.com/hasemeneh/golang-simple-login/src/repositories"
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
