package db

import (
	"../bash_profile"
	"../models/request"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (odb *MongoDBClient) WebRegister(user request.WebRegisterData) (request.WebUser, error) {
	collection := Use(odb.Client, bash_profile.DBName, bash_profile.WebUserCollection)
	var result request.WebUser
	result.Email = user.Email
	result.Username = user.Username
	result.Password = user.Password
	_id , err := collection.InsertOne(GetContext(), result)
	result.ID = _id.InsertedID
	return result, err
}

func (odb *MongoDBClient) WebLogin(data request.WebLoginData) (request.WebLoginResponseData, error) {
	var user request.WebLoginResponseData
	var users []request.WebLoginResponseData
	collection := Use(odb.Client, bash_profile.DBName, bash_profile.WebUserCollection)
	rst, _ := collection.Find(GetContext(), bson.D{{}})
	for rst.Next(context.TODO()) {
		var elem request.WebLoginResponseData
		err := rst.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, elem)
	}
	log.Print(users)
	result := collection.FindOne(GetContext(), bson.D{{"email", data.Email}}).Decode(&user)

	if result != nil {
		return user, result
	}

	return user, nil
}
