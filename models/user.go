package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/kyp-structs"
	"golang.org/x/crypto/bcrypt"
)

type User structs.User

func (repo *DB) CreateUser(u *User) *gorm.DB {
	password := u.Password
	u.Password = ""
	u.EncryptedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return repo.Create(u)
}

func (repo *DB) FindUserById(u *User, id int) *gorm.DB {
	return repo.Find(&u, id)
}

func (repo *DB) FindUserByEmail(u *User, email string) *gorm.DB {
	return repo.Where("email = ?", email).Find(&u)
}

func (u *User) VerifyPassword(password string) (bool, error) {
	return bcrypt.CompareHashAndPassword(u.EncryptedPassword, []byte(password)) == nil, nil
}
