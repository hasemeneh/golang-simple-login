package user

import (
	"context"

	"github.com/hasemeneh/golang-simple-login/src/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *userDomain {
	return &userDomain{
		DB: DB,
	}
}

const RegisterQuery = "INSERT INTO `user` (`email`, `username`, `address`, `password`) VALUES (?, ?, ?, ?);"

func (u *userDomain) InsertUser(ctx context.Context, userModel *models.UserModel) error {
	hasehedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.PlainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = u.DB.ExecContext(ctx, RegisterQuery, userModel.Email, userModel.Username, userModel.Address, hasehedPassword)
	return err
}

const updateAddressQuery = "UPDATE `user` SET `address` = ? WHERE `user`.`id` = ?;"

func (u *userDomain) UpdateAddress(ctx context.Context, address string, UserID int64) error {
	_, err := u.DB.ExecContext(ctx, RegisterQuery, address, UserID)
	return err
}

const updatePasswordQuery = "UPDATE `user` SET `password` = ? WHERE `user`.`id` = ?;"

func (u *userDomain) UpdatePassword(ctx context.Context, password string, UserID int64) error {
	hasehedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = u.DB.ExecContext(ctx, RegisterQuery, hasehedPassword, UserID)
	return err

}

const getUserByUsernameQuery = "SELECT id, email, username, address, password  FROM `user` WHERE `username` = ?"

func (u *userDomain) GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error) {
	resp := models.UserModel{}
	err := u.DB.GetContext(ctx, &resp, getUserByUsernameQuery, username)
	return &resp, err
}

const getUserByEmailQuery = "SELECT id  FROM `user` WHERE `email` = ?"

func (u *userDomain) GetUserByEmail(ctx context.Context, email string) (*models.UserModel, error) {
	resp := models.UserModel{}
	err := u.DB.GetContext(ctx, &resp, getUserByEmailQuery, email)
	return &resp, err
}

const getUserByIDQuery = "SELECT id, email, username, address, password  FROM `user` WHERE `id` = ?"

func (u *userDomain) GetUserByID(ctx context.Context, id int64) (*models.UserModel, error) {
	resp := models.UserModel{}
	err := u.DB.GetContext(ctx, &resp, getUserByIDQuery, id)
	return &resp, err
}
