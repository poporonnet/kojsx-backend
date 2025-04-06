package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type Data struct {
	id       id.SnowFlakeID
	name     string
	email    string
	password string
	role     int
}

func NewData(id id.SnowFlakeID, name string, email string, password string, role int) *Data {
	return &Data{id: id, name: name, email: email, password: password, role: role}
}

func (d Data) GetID() id.SnowFlakeID {
	return d.id
}

func (d Data) GetName() string {
	return d.name
}

func (d Data) GetEmail() string {
	return d.email
}

func (d Data) GetPassword() string {
	return d.password
}

// IsAdmin role -> 0
func (d Data) IsAdmin() bool {
	return d.role == 0
}

// IsVerified role -> 0,1
func (d Data) IsVerified() bool {
	return d.role != 2
}

// ToDomain DTOをドメインモデルに変換
func (d Data) ToDomain() model.User {
	u, _ := model.NewUser(d.GetID(), d.GetName(), d.GetEmail())
	if d.IsVerified() {
		u.SetVerified()
	}
	if d.IsAdmin() {
		u.SetAdmin()
	}

	u.SetPassword(d.GetPassword())
	return *u
}

// DomainToData ドメインモデルをDTOに変換
func DomainToData(in model.User) Data {
	role := 1
	if !in.IsVerified() {
		role = 2
	}
	if in.IsAdmin() {
		role = 0
	}
	return *NewData(in.GetID(), in.GetName(), in.GetEmail(), in.GetPassword(), role)
}
