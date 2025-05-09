package models

type OtpService struct {
	Email		string		`json:"email"`
}

type OtpVerify struct {
	Email		string		`json:"email"`
	Otp			string		`json:"otp"`
}
