package handler

import (
	"../../bash_profile"
	"../../models/request"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (h *Handler) WxLogin(id string) (request.WxUser, error)  {
	user, err := h.db.WxLogin(id)
	return user, err
}

func (h *Handler) WxRegister(c *gin.Context) {
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

	var wxRequestData request.WxRequestData

	err :=  c.ShouldBindJSON(&wxRequestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"error": err.Error(),
			"msg": "请求数据错误",
		})
		c.Abort()
		return
	}

	// 拼接微信API地址
	url := bash_profile.WxSite + "appid=" + bash_profile.WxAppId + "&secret=" + bash_profile.WxSecret + "&js_code=" + wxRequestData.Code + bash_profile.WxHttpTail

	// 转发请求到微信接口
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code": 503,
			"error": err.Error(),
			"msg": "第三方服务错误",
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
			"msg": "第三方数据读取失败",
			"data": nil,
		})
		c.Abort()
		return
	}

	respData := new(request.Code2Session)

	// 解码数据并赋值给返回的 data
	json.Unmarshal(body, &respData)

	// 检测openid是否为空
	if len(respData.Openid) == 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code": respData.Errcode,
			"error": respData.Errmsg,
			"msg": "第三方服务错误",
			"data": nil,
		})
		c.Abort()
		return
	}

	/**
	如果不进行数据库操作请注释掉下方数据库相关代码
	 */
	// mongodb
	/**
	避免同一账号重复写入数据库的两种方法:
	1. 使用mongo-driver 去数据库中查询openID是否存在，
	   如果存在则不进行写入操作
	2. 通过shell 命令手动设置openID为索引字段，
	   此时如果插入重复的openID 则会返回error

	推荐使用方法二，可以避免查询数据库，减少请求时间
	这里给出方法一的代码
	  */

	// 进行查询数据库操作
	userData, err := h.db.WxLogin(respData.Openid)
	log.Print(userData, err)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"error": nil,
			"msg": "登录成功",
			"Data": userData,
		})

		c.Abort()
		return
	}


	var user request.WxUser
	user.OpenID = respData.Openid
	// user.UnionID = respData.Unionid 如果使用union ID
	user.AvatarUrl = wxRequestData.AvatarUrl
	user.Nickname = wxRequestData.NickName
	user.City = wxRequestData.City
	user.Province = wxRequestData.Province
	user.Country = wxRequestData.Country
	user.Gender = wxRequestData.Gender
	user.CreateDate = time.Now()


	userData, err = h.db.WxRegister(user)

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