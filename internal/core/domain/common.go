package domain

const (
	SystemUUID = "00000000-0000-0000-0000-000000000000"
)

type HttpErrorCode string

type UserSearhKey string

var (
	Email UserSearhKey = "Email"
	Name  UserSearhKey = "Name"
)

var (
	InternalServerError HttpErrorCode = "500_InternalServerError"
	BadRequestError     HttpErrorCode = "401_BadRequestError"
)
