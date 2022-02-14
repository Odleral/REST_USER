package model

import (
	"github.com/satori/go.uuid"
	_ "gorm.io/gorm"
	"time" //nolint:gci
)

type User struct {
	UUID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Firstname string
	Lastname  string
	Email     string
	Age       uint
	CreatedAt time.Time
}

type Users []*User

type UserForm struct {
	Firstname string `json:"firstName" form:"required,max=255"`
	Lastname  string `json:"lastName" form:"required,max=255"`
	Email     string `json:"email" form:"required,max=255"`
	Age       uint   `json:"age" form:"required"`
}

func (f *UserForm) ToModel() (*User, error) {
	id := uuid.NewV4()

	return &User{
		UUID:      id,
		Firstname: f.Firstname,
		Lastname:  f.Lastname,
		Email:     f.Email,
		Age:       f.Age,
		CreatedAt: time.Now(),
	}, nil
}

type UserDto struct {
	UUID      uuid.UUID `json:"uuid"`
	Firstname string    `json:"firstName"`
	Lastname  string    `json:"lastName"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"create_at"`
}

type UsersDto []*UserDto

func (u *User) ToDto() *UserDto {
	return &UserDto{
		UUID:      u.UUID,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Age:       u.Age, //nolint:govet
		CreatedAt: u.CreatedAt,
	}
}

func (bs Users) ToDtos() UsersDto {
	dtos := make([]*UserDto, len(bs))
	for i, u := range bs {
		dtos[i] = u.ToDto()
	}

	return dtos
}
