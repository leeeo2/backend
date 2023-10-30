package domain

// register
type RegisterInput struct {
	Name      string
	Telephone string
	Password  string
}

type RegisterOutput struct {
	UserId uint
}

// login
type LoginInput struct {
	Telephone string
	Password  string
}

type LoginOutput struct {
	Token string
}

// get user info
type GetUserInfoInput struct {
	UserId    uint
	Telephone string
}

type GetUserInfoOutput struct {
	UserId    uint
	Name      string
	Telephone string
}
