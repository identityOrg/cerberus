package config

import "github.com/spf13/viper"

type SecretConfig struct {
	SessionSecret string
	TokenSecret   string
}

func NewSecretConfig() *SecretConfig {
	sec := &SecretConfig{
		SessionSecret: "jdhfbwjhebajwhevbahwevbahevbajwhevblawev",
		TokenSecret:   "jhkawhjebawhebvajebvkjahebvkjahebvkjebvj",
	}
	_ = viper.UnmarshalKey("secret", sec)
	return sec
}
