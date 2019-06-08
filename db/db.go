package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)



// 使用mongoDB官方库进行数据库操作
type MongoDBClient struct {
	*mongo.Client
}

// 建立与mongodb的连接
func NewConnection(uri string) (*MongoDBClient, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	return &MongoDBClient{
		client,
	}, err
}

// 选择数据库及表单
func Use(client *mongo.Client, dbname, collname string) ( *mongo.Collection ) {
	return client.Database(dbname).Collection(collname)
}

// 限制每次操作数据库的操作时间为10s
func GetContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

