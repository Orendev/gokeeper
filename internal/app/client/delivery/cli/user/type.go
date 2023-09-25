package user

type RegisterUserArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginUserArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
