package models

import (
	"time"

	"gopkg.in/hlandau/passlib.v1"

	. "massliking/backend/errors"
)

type User struct {
	Id        int64     `json:"id"         xorm:"pk autoincr index"`
	Username  string    `json:"username"   xorm:"notnull index unique"`
	Password  string    `json:"-"          xorm:"varchar(200) notnull"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func (u *User) Verify(password string) bool {
	_, err := passlib.Verify(password, u.Password)
	if err != nil {
		return false
	}

	return true
}

func CreateUser(c *Credentials) (*User, error) {
	user := &User{}

	hash, err := passlib.Hash(c.Password)
	if err != nil {
		return user, On(err, MODEL_USER_NOT_CREATED)
	}

	user.Username = c.Username
	user.Password = hash

	_, err = Engine.Insert(user)

	return user, On(err, MODEL_USER_NOT_CREATED)
}

func GetUser(username string) (*User, error) {
	user := &User{}

	_, err := Engine.
		Where("username = ?", username).
		Get(user)

	return user, On(err, MODEL_USER_NOT_FOUND)
}
