package handler

import (
	"../../models/request"
	myjwt "../jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) WebRegister(c *gin.Context)  {
	// 保证数据库接口已经被初始化
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"error": nil,
			"msg": "数据库连接失败",
		})
		c.Abort()
		return
	}


	var webRegisterData request.WebRegisterData

	err :=  c.ShouldBindJSON(&webRegisterData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"error": "",
			"msg": "请求数据错误",
		})
		c.Abort()
		return
	}
	log.Print(webRegisterData)
	userData, err := h.db.WebRegister(webRegisterData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"error": err.Error(),
			"msg": "注册失败",
			"data": nil,
		})
		c.Abort()
		return
	}

	// 返回数据给前台完成register
	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"error": nil,
		"msg": "注册成功",
		"Data": userData,
	})

	return
}

func (h Handler) WebLogin(c *gin.Context)  {
	// 保证数据库接口已经被初始化
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"error": nil,
			"msg": "数据库连接失败",
		})
		c.Abort()
		return
	}


	var webLoginData request.WebLoginData
	err :=  c.ShouldBindJSON(&webLoginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"error": err.Error(),
			"msg": "请求数据错误",
		})
		c.Abort()
		return
	}


	log.Print(webLoginData)
	userData, err := h.db.WebLogin(webLoginData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"error": err.Error(),
			"msg": "登录失败",
			"data": nil,
		})
		c.Abort()
		return
	}

	claims := myjwt.JWTClaims{
		ID: userData.ID,
		Email: userData.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer: "admin",	// 签名的发行者
		},
	}

	token, jwtErr := myjwt.CreateToken(claims)
	if jwtErr  != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"error": jwtErr,
			"msg": "无法生成token",
			"token": nil,
		})
	}

	// 返回数据给前台完成register
	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"error": nil,
		"msg": "登录成功",
		"token": token,
	})

	return
}

func  TestJwt(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.JWTClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "测试成功",
			"data": claims,
		})
	}
}