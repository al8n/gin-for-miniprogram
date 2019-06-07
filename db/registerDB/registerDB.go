package registerDB

import (
	"../../bash_profile"
	"../../db"
	"../../models/register"
	"time"
)

type Message struct {
	Msg string
	Code int
	Error string
}

type _RegisterDBInterface interface {
	WxRegister(user register.WxUser) Message
}

type UserDB struct {
	dbIpt _RegisterDBInterface
}

func (dbIpt *UserDB) WxRegister(user register.WxUser) Message {
	// 连接数据库
	collection := db.NewConnection(bash_profile.DBConnect).Use(bash_profile.DBName, bash_profile.UserCollection)

	// 用户创建时间
	user.CreateDate = time.Now()
	user.UpdateDate = time.Now()

	// 插入数据库
	// 第一项返回值为数据库自动生成的ObjectId
	// 如需要获取此ID并返回给前台请自行处理
	_, err := collection.InsertOne(db.GetContext(), user)

	if err != nil {
		return Message{
			Error: err.Error(),
			Code: 0,
			Msg: "Internal error",
		}
	}

	return Message{
		Code: 1,
		Error: "",
		Msg: "User added",
	}
}

