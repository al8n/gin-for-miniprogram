package wxHandler

import (
	"../../bash_profile"
	"../../db/registerDB"
	"../../models/register"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type _WxHandlerInterface interface {
	WxRegister(c *gin.Context)
	WxLogin(c *gin.Context)
}

type WxHandler struct {
	wxDB _WxHandlerInterface
}



//type wxOpenID struct {
//	openid string `json:"wxOpenID"`
//	// unionid string `json:"wxUionID"` // 如果使用union ID 请注释掉上一行代码并使用这行代码
//}

var wx *registerDB.UserDB

func (wxDB *WxHandler) WxRegister(c *gin.Context)  {
	var data register.WxRegister

	err :=  c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"error": err.Error(),
			"msg": "Form information format error",
		})
		c.Abort()
		return
	}
	// 拼接微信API地址
	url := bash_profile.WxSite + "appid=" + bash_profile.WxAppId + "&secret=" + bash_profile.WxSecret + "&js_code=" + data.Code + bash_profile.WxHttpTail

	// 转发请求到微信接口
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code": 503,
			"error": err.Error(),
			"msg": "Third-party service error",
			"data": nil,
		})
		c.Abort()
		return
	}

	defer resp.Body.Close()

	// 读取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"error": err.Error(),
			"msg": "Data read failed",
			"data": nil,
		})
		c.Abort()
		return
	}

	respData := new(register.Code2Session)

	// 解码数据并赋值给返回的 data
	json.Unmarshal(body, &respData)

	/**
	如果不进行数据库操作请注释掉下方数据库相关代码
	 */
	 // mongodb
	var user register.WxUser
	user.OpenID = respData.Openid
	// user.UnionID = respData.Unionid 如果使用union ID
	user.AvatarUrl = data.AvatarUrl
	user.Nickname = data.NickName
	user.City = data.City
	user.Province = data.Province
	user.Country = data.Country
	user.Gender = data.Gender

	dbData := wx.WxRegister(user)
	if dbData.Code == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"error": dbData.Error,
			"msg": dbData.Msg,
			"data": nil,
		})
		c.Abort()
		return
	}


	// 返回数据给前台完成register
	c.JSON(http.StatusCreated, gin.H{
	 	"code": 201,
	 	"error": nil,
	 	"msg": dbData.Msg,
	 	"data": respData,
	})
	return
}