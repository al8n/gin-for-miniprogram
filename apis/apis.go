package apis

import (

	"./handler"
	"./jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {

	h, err := handler.NewHandler()

	if err != nil {
		return err
	}

	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h handler.IHandler) error {

	r := gin.Default()

	// 跨域
	r.Use(cors.Default())

	// 微信小程序api
	wxGroup := r.Group("/wx")
	{
		wxGroup.POST("/register", h.WxRegister)
	}

	// web 端API
	webGroup := r.Group("/web")
	{
		webGroup.POST("/register", h.WebRegister)
		webGroup.POST("/login", h.WebLogin)
	}

	testGroup := r.Group("/test", jwt.JWTAuth())
	{
		testGroup.POST("", handler.TestJwt)
	}
	// 如果需读取静态文件
	r.Static("/imgs", "../assets/imgs")
	r.Static("/videos", "../assets/videos")

	return r.Run(address)
}

