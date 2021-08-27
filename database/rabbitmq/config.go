package rabbitmq

type RabbitMQConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	Vhost    string
	Exchange string
}

func defaultDB() *RabbitMQConfig {
	dbCfg := &RabbitMQConfig{}
	dbCfg.Host = "localhost"
	dbCfg.Port = "5672"
	dbCfg.Vhost = "/"
	dbCfg.Exchange = "thetan"
	return dbCfg
}
