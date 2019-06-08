package request

import (
	"time"
)

type WxUser struct {
	ID			interface{}			`bson:"_id,omitempty"`	// mongodb创建时自动生成的ID
	OpenID      string  			`bson:"wxOpenId" binding:"required"`
	// UnionID 	string				`bson:"wxUnionId"` 如果程序使用unionid 登录
	Nickname    string 				`bson:"wxNickname"`
	AvatarUrl   string 				`bson:"wxAvatarUrl"`
	Gender		int 				`bson:"gender"`
	City 		string				`bson:"city"`
	Province	string				`bson:"province"`
	Country 	string				`bson:"country"`
	CreateDate	time.Time  			`bson:"createDate"`
	/**
	put custom code here
	eg.
		Email			string		`json:"email" binding:"required"`
		Password 		string		`json:"password" binding:"required"`
		UpdateDate  time.Time           `bson:"updated"`
	 */
}


type WxRequestData struct {
	Code		string      `json:"code"`
	NickName	string 		`json:"nickName"`
	Gender		int 		`json:"gender"`
	City 		string		`json:"city"`
	Province	string		`json:"province"`
	Country 	string		`json:"country"`
	AvatarUrl	string 		`json:"avatarUrl"`
	/**
	put custom code here
	eg.
		Email			string		`json:"email" binding:"required"`
		Password 		string		`json:"password" binding:"required"`
	 */
}

type Code2Session struct {
	Errcode    int32  `json:"errcode"`		// 错误码
	Errmsg     string `json:"errmsg"`		// 错误信息
	Openid     string `json:"openid"` 		// 用户唯一标识
	// Unionid    string `json:"unionid"` 	// 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见微信 UnionID
	SessionKey string `json:"session_key"`	// 会话密钥
}