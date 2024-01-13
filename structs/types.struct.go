package structs

type TOTP struct {
	Otp        string
	Otp_Expiry string
}

type File struct {
	FileBytes     []byte
	FileName      string
	FileExtension string
}
