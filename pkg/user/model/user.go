package model

import (
	"unicode/utf8"

	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type UserRole int

const (
	Admin = iota
	Normal
	Unverified
)

// User ユーザー
type User struct {
	id       id.SnowFlakeID
	name     string
	email    string
	password string
	role     UserRole
}

// UserNameLengthError ユーザー名の長さエラー
type UserNameLengthError struct {
}

func (e UserNameLengthError) Error() string {
	return "ユーザー名の長さが不正です"
}

// UserEmailLengthError メールアドレスの長さエラー
type UserEmailLengthError struct {
}

func (e UserEmailLengthError) Error() string {
	return "メールアドレスの長さが不正です"
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

// IsVerified role -> NOT 2
func (u *User) IsVerified() bool {
	return u.role == Unverified
}

// IsAdmin role -> 0
func (u *User) IsAdmin() bool {
	return u.role == Admin
}

func (u *User) SetAdmin() {
	if !u.IsAdmin() {
		u.role = Admin
	}
}

func (u *User) SetNormal() {
	if u.IsAdmin() {
		u.role = Normal
	}
}

func (u *User) SetVerified() {
	if !u.IsVerified() {
		u.role = Normal
	}
}

func (u *User) SetPassword(p string) {
	u.password = p
}
