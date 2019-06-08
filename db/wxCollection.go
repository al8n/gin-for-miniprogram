package db

import (
	"../bash_profile"
	"../models/request"
	"go.mongodb.org/mongo-driver/bson"
)

// 实现数据库操作
func (odb *MongoDBClient) WxRegister(user request.WxUser) (request.WxUser, error) {
	collection := Use(odb.Client, bash_profile.DBName, bash_profile.WxUserCollection)
	_id , err := collection.InsertOne(GetContext(), user)
	user.ID = _id.InsertedID
	return user, err
}

func (odb *MongoDBClient) WxLogin(id string) (request.WxUser, error) {
	var user request.WxUser
	collection := Use(odb.Client, bash_profile.DBName, bash_profile.WxUserCollection)
	result := collection.FindOne(GetContext(), bson.D{{"wxOpenId", id}}).Decode(&user)

	if result != nil {
		return user, result
	}

	return user, nil
}