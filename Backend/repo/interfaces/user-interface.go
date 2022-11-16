package interfaces


type User struct{
	UserName string `json:"userName"`
    UserEmail string `json:"userEmail"`
	UserPhone string `json:"userPhone"`
	Profile string 	`json:"profile"`
}


type UserSourceInterface interface{
	CreateUser(userDetails User)
}

