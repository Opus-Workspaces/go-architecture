package config

type BaseDatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Timeout      int
}

type Redis struct {
	Host        string
	Port        string
	Password    string
	DB          int
	DefaultDB   int
	MinIdleCons int
	PoolSize    int
	IdleTimeout int
}

type DatabaseType struct {
	MongoDB BaseDatabaseConfig
	Redis   *Redis
}
