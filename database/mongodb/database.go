package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDB() *mongo.Database {
	return db
}

type DBConfig struct {
	DbName     string
	UserName   string
	Password   string
	Host       string
	Port       string
	IsReplica  bool
	ReplicaSet string
}

// MongoConfig new version
type MongoConfig struct {
	DbName            string
	ConnectionUrl     string
	MaxConnectionPool uint64
}

func defaultDB() *DBConfig {
	dbCfg := &DBConfig{}
	dbCfg.Host = "localhost"
	dbCfg.Port = "27017"
	dbCfg.DbName = "db_default"
	return dbCfg
}

func ConnectMongoWithConfig(dbConfig *MongoConfig, conf *Config) (context.Context, *mongo.Client, context.CancelFunc) {
	if conf == nil {
		conf = defaultConf()
	}
	if dbConfig == nil {
		dbConfig = &MongoConfig{
			DbName:        "thetan",
			ConnectionUrl: "mongodb://127.0.0.1:27017",
		}
	}
	config = conf
	dbName = dbConfig.DbName

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOption := options.Client().ApplyURI(dbConfig.ConnectionUrl)
	if dbConfig.MaxConnectionPool > 0 {
		clientOption.SetMaxPoolSize(dbConfig.MaxConnectionPool)
	}

	clientNew, err := NewClient(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}
	client = clientNew

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("[FATAL] CAN'T CONNECTING TO MONGODB: %s", err.Error())
	}

	db = client.Database(dbName)

	log.Printf("[INFO] CONNECTED TO MONGO DB %s", dbName)
	return ctx, client, cancel
}

func SetDefaultConfig(dbConfig *DBConfig, conf *Config) (context.Context, *mongo.Client, context.CancelFunc) {
	if conf == nil {
		conf = defaultConf()
	}
	if dbConfig == nil {
		dbConfig = defaultDB()
	}

	config = conf
	dbName = dbConfig.DbName
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientNew, err := NewClient(ctx, options.Client().ApplyURI(buildUri(dbConfig)))
	if err != nil {
		panic(err)
	}
	client = clientNew
	db = client.Database(dbName)

	log.Printf("[INFO] CONNECTED TO MONGO DB %s", dbName)
	return ctx, client, cancel
}

func buildUri(dbConfig *DBConfig) string {
	username := dbConfig.UserName
	password := dbConfig.Password
	host := dbConfig.Host
	port := dbConfig.Port

	link := fmt.Sprintf("%s:%s/?w=majority", host, port)
	if dbConfig.IsReplica {
		link = fmt.Sprintf("%s", dbConfig.ReplicaSet)
	}
	var uri string
	if username == "" && password == "" {
		uri = fmt.Sprintf("mongodb://%s", link)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s", username, password, link)
	}
	log.Println("MongoDb buildUri = ", uri)
	return uri
}
