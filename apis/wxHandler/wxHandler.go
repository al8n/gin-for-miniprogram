package wxHandler

import (
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

const (
	site = "https://api.weixin.qq.com/sns/jscode2session?"
	wxAppId = "wx6f9cc63ed9ded89b" // replaced by your App id
	wxSecret = "421169661fc3ef99110fe20c75a64ad1" // replaced by your secret
	httpTail = "&grant_type=authorization_code"
)

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
	url := site + "appid=" + wxAppId + "&secret=" + wxSecret + "&js_code=" + data.Code + httpTail

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