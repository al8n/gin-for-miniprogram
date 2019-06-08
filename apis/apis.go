package apis

import (
	"./Handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {

	r := gin.Default()

	// 跨域
	r.Use(cors.Default())

	h, _ := handler.NewHandler()
	// 微信小程序api
	wxGroup := r.Group("/wx")
	{
		wxGroup.POST("/register", h.WxRegister)
	}

	return r.Run(address)
}
