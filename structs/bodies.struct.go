package structs

type ForgotPassword struct {
	Email string `json:"email"`
}

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Login struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
