package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)



// 使用mongoDB官方库进行数据库操作
type _MongoDBClient struct {
	cli *mongo.Client
}

// 建立与mongodb的连接
func NewConnection(uri string) (conn *_MongoDBClient) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// 5s 之内如果无法连接到数据库则抛出异常
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	conn = &_MongoDBClient{client}
	return conn
}

// 选择数据库及表单
func (conn *_MongoDBClient) Use(dbname, collname string) (collection *mongo.Collection) {
	return conn.cli.Database(dbname).Collection(collname)
}

func GetContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}