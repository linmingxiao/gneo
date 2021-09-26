package mongo

import (
	"context"
	"github.com/linmingxiao/gneo/logx"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	_ "go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

// Redigo初始化的配置参数
type (
	ConnConfig struct {
		Uri      string `json:",optional"`
		Database string `json:",optional"`
		MasterName string `json:",optional"`
		MaxPoolSize uint64 `json:",optional"`
	}
	MgoX struct {
		DB *mongo.Database
		Cli *mongo.Client
		Ctx context.Context
	}
)

func (mgX *MgoX) Close() {
	err := mgX.Cli.Disconnect(mgX.Ctx)
	if err != nil {
		logx.Error(err)
	}
	logx.Infof("MongoDB %v closed.", mgX.Cli)
}

func NewMongo(cf *ConnConfig) *MgoX {
	uri := cf.Uri
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientUri := options.Client().ApplyURI(uri).SetMaxPoolSize(cf.MaxPoolSize)
	client, err := mongo.Connect(ctx, clientUri)
	if err != nil {
		logx.Errorf("MongoDB connect err: %v", err)
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logx.Errorf("MongoDB ping err: %v", err)
		panic(err)
	}
	dataBase := client.Database(cf.Database)
	logx.Infof("MongoDB: %s connect sucess!", cf.MasterName)
	return &MgoX{DB: dataBase, Cli: client, Ctx: context.Background()}
}
