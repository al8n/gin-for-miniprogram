package db

import (
	"../models/request"
)

type ODBLayer interface {
	WxRegister(request.WxUser) (request.WxUser, error)
	WxLogin(string) (request.WxUser, error)
}