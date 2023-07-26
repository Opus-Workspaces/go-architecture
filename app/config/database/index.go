package config

import (
	cfg "go-architecture/app/config/consts"

	"github.com/spf13/viper"
)

func ConfigDatabase() DatabaseType {
	var mongo = cfg.PrefixDB + ".mongo"
	var redis = cfg.PrefixDB + ".redis"

	mongoHost := viper.GetString(mongo + ".host")
	mongoPort := viper.GetInt(mongo + ".port")
	mongoUser := viper.GetString(mongo + ".user")
	mongoPassword := viper.GetString(mongo + ".password")
	mongoDatabaseName := viper.GetString(mongo + ".db_name")
	mongoTimeout := viper.GetInt(mongo + ".timeout")

	redisHost := viper.GetString(redis + ".host")
	redisPort := viper.GetString(redis + ".port")
	redisPassword := viper.GetString(redis + ".password")
	redisDB := viper.GetInt(redis + ".db")
	redisDefaultDB := viper.GetInt(redis + ".default_db")
	redisMinIdleCons := viper.GetInt(redis + ".min_idle_cons")
	redisPoolSize := viper.GetInt(redis + ".pool_size")
	redisIdleTimeout := viper.GetInt(redis + ".idle_timeout")

	return DatabaseType{
		MongoDB: BaseDatabaseConfig{
			Host:         mongoHost,
			Port:         mongoPort,
			User:         mongoUser,
			Password:     mongoPassword,
			DatabaseName: mongoDatabaseName,
			Timeout:      mongoTimeout,
		},
		Redis: &Redis{
			Host:        redisHost,
			Port:        redisPort,
			Password:    redisPassword,
			DB:          redisDB,
			DefaultDB:   redisDefaultDB,
			MinIdleCons: redisMinIdleCons,
			PoolSize:    redisPoolSize,
			IdleTimeout: redisIdleTimeout,
		},
	}
}
