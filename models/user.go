package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/kyp-structs"
	"github.com/rafaeljesus/kyp-users/db"
	"golang.org/x/crypto/bcrypt"
)

type User structs.User

func (u *User) Create() *gorm.DB {
	password := u.Password
	u.Password = ""
	u.EncryptedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return db.Repo.Create(u)
}

func (u *User) FindById(id int) *gorm.DB {
	return db.Repo.Find(&u, id)
}

func (u *User) FindByEmail(email string) *gorm.DB {
	return db.Repo.Where("email = ?", email).Find(&u)
}

func (u *User) VerifyPassword(password string) (bool, error) {
	return bcrypt.CompareHashAndPassword(u.EncryptedPassword, []byte(password)) == nil, nil
}
