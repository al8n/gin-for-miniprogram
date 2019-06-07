package apis

import (
	"./wxHandler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {

	r := gin.Default()

	// 跨域
	r.Use(cors.Default())

	// 微信小程序api
	wxGroup := r.Group("/wx")
	{
		wxController := new(wxHandler.WxHandler)
		// wxGroup.POST("/login", wxController.WxLogin)
		wxGroup.POST("/register", wxController.WxRegister)
	}

	return r.Run(address)
}
