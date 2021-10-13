package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetDB() *mongo.Database {
	return db
}

type DBConfig struct {
	DbName   string
	UserName string
	Password string
	Host  string
	Port   string
	IsReplica bool
	ReplicaSet string
}
func defaultDB() *DBConfig {
	dbCfg := &DBConfig{}
	dbCfg.Host = "localhost"
	dbCfg.Port = "27017"
	dbCfg.DbName = "db_default"
	return dbCfg
}

func SetDefaultConfig(dbConfig *DBConfig, conf *Config) (context.Context, *mongo.Client, context.CancelFunc) {
	if conf == nil {
		conf = defaultConf()
	}
	if dbConfig == nil {
		dbConfig = defaultDB()
	}

	config = conf
	dbName := dbConfig.DbName
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := NewClient(ctx, options.Client().ApplyURI(buildUri(dbConfig)))
	if err != nil {
		panic(err)
	}
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
		link = fmt.Sprintf("%s/theta?ssl=false", dbConfig.ReplicaSet)
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
