package dao

import (
	"github.com/leeeo2/backend/pkg/model"
)

func CreateUser(user *model.User) (uint, error) {
	tx := model.GetDB()
	tx = tx.Create(user)
	return user.ID, tx.Error
}

func DeleteUserById(id uint) error {
	tx := model.GetDB()
	tx = tx.Where("id = ?", id)
	return tx.Delete(&model.User{}).Error
}

func DeleteUserByTelephone(telephone string) error {
	tx := model.GetDB()
	tx = tx.Where("telephone = ?", telephone)
	return tx.Delete(&model.User{}).Error
}

func DescribeUserById(id uint) (*model.User, error) {
	tx := model.GetDB()
	var user model.User
	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DescribeUserByTelephone(telephone string) (*model.User, error) {
	tx := model.GetDB()
	var user model.User
	err := tx.Where("telephone = ?", telephone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DescribeUsers(input *DescribeInput) ([]*model.User, error) {
	tx := GetDescribeTx(model.GetDB(), input)

	users := make([]*model.User, 0)
	err := tx.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func IsUserExist(telephone string) bool {
	tx := model.GetDB()
	var user model.User
	tx.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
