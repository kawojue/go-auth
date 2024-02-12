package structs

import "mime/multipart"

type ForgotPassword struct {
	Email string `json:"email"`
}

type ResetPassword struct {
	Otp   string `json:"otp"`
	Pswd  string `json:"pswd"`
	Pswd2 string `json:"pswd2"`
}

type SignUp struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Username  string `json:"username"`
}

type Login struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type AvatarForm struct {
	Avatar *multipart.FileHeader `json:"avatar"`
}
