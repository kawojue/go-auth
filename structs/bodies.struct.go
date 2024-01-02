package structs

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Login struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
