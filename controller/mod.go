package controller

import (
	"common-web-framework/middle"
	"common-web-framework/models"
)

func GetUserIdIfSuper(user models.User) int {
	if user.Role.Name == string(middle.SuperRole) {
		return -1
	}

	return user.ID
}
