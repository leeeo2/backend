package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"varchar(255);not null"`
}

func encryptPassword(passwd string) (string, error) {
	if passwd == "" {
		return "", nil
	}
	if hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	hash, err := encryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) error {
	if u.Password != "" {
		hash, err := encryptPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hash
	}
	return nil
}
