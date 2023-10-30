package dto

import "github.com/leeeo2/backend/pkg/model"

type UserInfoDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserInfoDto(user *model.User) *UserInfoDto {
	return &UserInfoDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
