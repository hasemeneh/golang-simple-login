package service

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hasemeneh/golang-simple-login/helper/redis"
	"github.com/hasemeneh/golang-simple-login/src/constants"
	"github.com/hasemeneh/golang-simple-login/src/definitions"
	user_domain "github.com/hasemeneh/golang-simple-login/src/domain/user"
	user_usecase "github.com/hasemeneh/golang-simple-login/src/usecase/user"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	UserUsecase definitions.UserUsecase
}

func New() *Service {
	db, err := sqlx.Open(constants.MySQLDatabaseDriver, "root:@tcp(localhost:3306)/simple_login?parseTime=true&loc=Local")

	if err != nil {
		log.Fatalln("error Opening Database", err)
	}
	redisInterface := redis.New("localhost:6379", "")
	userDomain := user_domain.New(db)
	userUsecase := user_usecase.New(userDomain, redisInterface)
	serviceObj := Service{
		UserUsecase: userUsecase,
	}
	return &serviceObj
}
