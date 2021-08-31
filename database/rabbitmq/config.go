package rabbitmq

type RabbitMQConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	Vhost    string
	Exchange string
	Durable  bool
}

func defaultDB() *RabbitMQConfig {
	dbCfg := &RabbitMQConfig{}
	dbCfg.Host = "localhost"
	dbCfg.Port = "5672"
	dbCfg.Vhost = "/"
	dbCfg.Exchange = "thetan"
	dbCfg.Durable = false
	return dbCfg
}
