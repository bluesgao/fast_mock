package dao

import (
	"context"
	"fast_mock/conf"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var mgoCli *MongoCli

type MongoCli struct {
	//uri    string //数据库网络地址(mongodb://ip:port)  (mongodb://username:password@ip:port)
	client *mongo.Client
}

func MongoClose() {
	mgoCli.client.Disconnect(context.TODO())
}

func GetMongoCli() *MongoCli {
	return mgoCli
}

func MongoInit(conf *conf.Conf) {
	log.Printf("database init conf:%+v \n", conf)
	//uri := "mongodb://admin:123456@47.97.205.190:27017"
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", conf.Mongo.Username, conf.Mongo.Password, conf.Mongo.Host, conf.Mongo.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //养成良好的习惯，在调用WithTimeout之后defer cancel()
	opts := &options.ClientOptions{}
	//opts.SetAuth(options.Credential{AuthMechanism: "SCRAM-SHA-1", Username: "admin", Password: "123456"})
	opts.SetMaxPoolSize(5) //设置连接池大小
	opts.ApplyURI(uri)
	log.Printf("mongo MongoInit opts:%+v \n", opts)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("mongo MongoInit err:%+v \n", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("mongo MongoInit Ping err:%+v \n", err)
	}

	mgoCli = &MongoCli{
		client: client,
	}
}

/**
	database   string //要连接的数据库
	collection string //要连接的集合
 */
func (m *MongoCli) GetCollection(database string, coll string) *mongo.Collection {
	collection := m.client.Database(database).Collection(coll)
	return collection
}

//插入单个
func (m *MongoCli) InsertOne(ctx context.Context, database string, coll string, value interface{}) (interface{}, error) {
	if ctx == nil {
		ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	}
	ret, err := m.GetCollection(database, coll).InsertOne(ctx, value)
	if err != nil {
		log.Printf("mongo InsertOne err:%+v \n", err)
		return nil, err
	}
	return ret.InsertedID, nil
}

// 查询单个
func (m *MongoCli) FindOne(ctx context.Context, database string, coll string, key string, value interface{}) (bson.M, error) {
	if ctx == nil {
		ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	}
	filter := bson.D{{key, value}}
	var result bson.M
	if err := m.GetCollection(database, coll).FindOne(ctx, filter).Decode(&result); err != nil {
		log.Printf("mongo FindOne err:%+v \n", err)
		return nil, err
	}
	return result, nil
}

// 查询多个
func (m *MongoCli) FindMany(ctx context.Context, database string, coll string, filter bson.M) ([]bson.M, error) {
	if ctx == nil {
		ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	}
	cursor, err := m.GetCollection(database, coll).Find(ctx, filter)
	if err != nil {
		log.Printf("mongo FindMany err:%+v \n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var resultArr []bson.M
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Printf("mongo FindMany Decode err:%+v \n", err)
			return nil, err
		}
		resultArr = append(resultArr, result)
	}
	return resultArr, nil
}
