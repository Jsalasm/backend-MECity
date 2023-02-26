package userController

import (
	userModel "go-project/model/user"
)

func GetAllUsers() []userModel.User {
	var Users = []userModel.User{
		{Id: "1", Name: "User1", Tel: "XXXX-XXXX", Email: "email@mail.com", Paswsord: "pwd"},
		{Id: "2", Name: "User2", Tel: "XXXX-XXXX", Email: "email@mail.com", Paswsord: "pwd"},
		{Id: "3", Name: "User3", Tel: "XXXX-XXXX", Email: "email@mail.com", Paswsord: "pwd"},
	}

	return Users
}
