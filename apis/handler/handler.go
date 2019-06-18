package handler

import (

	"../../bash_profile"
	"../../db"
	"../../models/request"
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	WxLogin(id string) (request.WxUser, error)
	WxRegister(c *gin.Context)
	WebRegister(c *gin.Context)
	WebLogin(c *gin.Context)
}

type Handler struct {
	db db.ODBLayer
}

func NewHandler() (IHandler, error)  {
	return NewHandlerWithParams()
}

func NewHandlerWithParams() (IHandler, error)  {
	client, err := db.NewConnection(bash_profile.DBConnect)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: client,
	}, nil
}

