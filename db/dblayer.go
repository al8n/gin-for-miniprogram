package db

import (
	"../models/request"
)

type ODBLayer interface {
	WxRegister(request.WxUser) (request.WxUser, error)
	WxLogin(string) (request.WxUser, error)

	WebRegister(user request.WebRegisterData) (request.WebUser, error)
	WebLogin(request.WebLoginData) (request.WebLoginResponseData, error)
}