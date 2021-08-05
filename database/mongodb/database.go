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
	var uri string
	if username == "" && password == "" {
		uri = fmt.Sprintf("mongodb://%s:%s/?w=majority", host, port)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/?w=majority", username, password, host, port)
	}
	return uri
}
