package config

import (
	cfg "go-architecture/app/config/consts"

	"github.com/spf13/viper"
)

func ConfigServer() ServerType {

	address := viper.GetString(cfg.PrefixHttpServer + ".address")
	port := viper.GetString(cfg.PrefixHttpServer + ".port")
	timeout := viper.GetUint16(cfg.PrefixHttpServer + ".timeout")
	appVersion := viper.GetString(cfg.PrefixHttpServer + ".app_version")
	readTimeout := viper.GetUint16(cfg.PrefixHttpServer + ".read_timeout")
	writeTimeout := viper.GetUint16(cfg.PrefixHttpServer + ".write_timeout")
	defaultTimeout := viper.GetUint16(cfg.PrefixHttpServer + ".default_timeout")

	return ServerType{
		Address:        address,
		Port:           port,
		Timeout:        timeout,
		AppVersion:     appVersion,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		DefaultTimeout: defaultTimeout,
	}
}
