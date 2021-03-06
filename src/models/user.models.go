package models

type UserModel struct {
	ID            int64  `db:"id" json:"id"`
	Email         string `db:"email" json:"email"`
	Username      string `db:"username" json:"username"`
	Address       string `db:"address" json:"address"`
	Password      []byte `db:"password" json:"-"`
	PlainPassword string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
