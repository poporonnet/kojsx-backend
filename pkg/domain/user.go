package domain

import (
	"unicode/utf8"

	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

// User ユーザー
type User struct {
	id       id.SnowFlakeID
	name     string
	email    string
	password string
	role     int
}

// UserNameLengthError ユーザー名の長さエラー
type UserNameLengthError struct {
	message string
}

func (e UserNameLengthError) Error() string {
	return e.message
}

// UserEmailLengthError メールアドレスの長さエラー
type UserEmailLengthError struct {
	message string
}

func (e UserEmailLengthError) Error() string {
	return e.message
}

/*
NewUser
不変: ID, Name, Email
*/
func NewUser(uID id.SnowFlakeID, name string, email string) (*User, error) {
	if utf8.RuneCountInString(name) > 64 || utf8.RuneCountInString(name) <= 0 {
		return nil, UserNameLengthError{}
	}

	// メールアドレスの最小文字数: <1>@<1>.<2> -> 6文字
	if utf8.RuneCountInString(email) > 64 || utf8.RuneCountInString(email) < 6 {
		return nil, UserEmailLengthError{}
	}

	return &User{
		id:    uID,
		name:  name,
		email: email,
		role:  2,
	}, nil
}

func (u *User) GetID() id.SnowFlakeID {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) IsVerified() bool {
	return u.role != 2
}

func (u *User) IsAdmin() bool {
	/*
		role:
			admin 0
			normal 1
			unverified 2
	*/
	return u.role == 0
}

func (u *User) SetAdmin() {
	if !u.IsAdmin() {
		u.role = 0
	}
}

func (u *User) SetNormal() {
	if u.IsAdmin() {
		u.role = 1
	}
}

func (u *User) SetVerified() {
	if !u.IsVerified() {
		u.role = 1
	}
}

func (u *User) SetPassword(p string) {
	u.password = p
}
