package users

type User struct {
	Id string	`json:"id"`
	Name string	`json:"name"`
	Tel string	`json:"tel"`
	Email string	`json:"email"`
	Paswsord string	`json:"password"`
}


func GetAllUsers () []User {
	var Users = []User{
		{Id: "1", Name: "User1", Tel: "XXXX-XXXX", Email: "email@mail.com", Paswsord: "pwd"},
		{Id: "2", Name: "User2", Tel: "XXXX-XXXX", Email: "email@mail.com", Paswsord: "pwd"},
		{Id: "3", Name: "User3", Tel: "XXXX-XXXX", Email: "email@mail.com", Paswsord: "pwd"},
	}

	return Users
}