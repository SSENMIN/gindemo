package dto

import "oceanlearn.teach/gindemo/model"

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User ) UserDto  {
	return UserDto{
		user.Name,
		user.Telephone,
	}
}