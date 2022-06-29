package cfg

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *viper.Viper

const (
	httpServerPort = "SERVER_PORT"
	dbHost         = "DB_HOST"
	dbPort         = "DB_PORT"
	dbPassword     = "DB_PASSWORD"
	dbName         = "DB_NAME"
	secretKey      = "SECRET_KEY"
)

func get() *viper.Viper {
	if cfg == nil {
		cfg = viper.New()

		cfg.BindEnv(httpServerPort)
		cfg.SetDefault(httpServerPort, ":8091")

		cfg.BindEnv(dbHost)
		cfg.SetDefault(dbHost, "localhost")
		cfg.BindEnv(dbPort)
		cfg.SetDefault(dbPort, "27017")
		cfg.BindEnv(dbName)
		cfg.SetDefault(dbName, "yas")
		cfg.BindEnv(dbPort)

		cfg.BindEnv(secretKey)
	}
	return cfg
}

func HttpServerPort() string {
	return get().GetString(httpServerPort)
}

func MongoDbURI() string {
	return fmt.Sprintf("mongodb://%s:%s", get().GetString(dbHost), get().GetString(dbPort))
}

func DbName() string {
	return get().GetString(dbName)
}

func SecretKey() []byte {
	return []byte(get().GetString(secretKey))
}
