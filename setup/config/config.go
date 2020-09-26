package config

import (
	"fmt"
	core "github.com/identityOrg/cerberus-core"
	"github.com/identityOrg/oidcsdk"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Driver string
	DSN    string
}

func NewDBConfig() *DBConfig {
	db := &DBConfig{}
	_ = viper.UnmarshalKey("db", db)
	return db
}

type ServerConfig struct {
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
	Demo  bool   `mapstructure:"demo"`
}

func NewServerConfig() *ServerConfig {
	s := &ServerConfig{}
	err := viper.UnmarshalKey("server", s)
	if err != nil {
		fmt.Printf("%v", err)
	}
	value := viper.GetBool("demo")
	if value {
		s.Demo = value
	}
	value = viper.GetBool("debug")
	if value {
		s.Debug = value
	}
	port := viper.GetString("addr")
	s.Port = port
	return s
}

type SecretConfig struct {
	SessionSecret string
	TokenSecret   string
}

func NewSecretConfig() *SecretConfig {
	sec := &SecretConfig{}
	_ = viper.UnmarshalKey("secret", sec)
	return sec
}

func NewSDKConfig() *oidcsdk.Config {
	oauth2Config := oidcsdk.NewConfig("http://localhost:8080")
	_ = viper.UnmarshalKey("provider", oauth2Config)
	return oauth2Config
}

func NewCoreConfig() *core.Config {
	coreConfig := &core.Config{}
	_ = viper.UnmarshalKey("core", coreConfig)
	return coreConfig
}
