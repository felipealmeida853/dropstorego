package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	MongoDBURI             string        `mapstructure: "MONGODBURI"`
	RedisURI               string        `mapstructure: "REDISURI"`
	Port                   string        `mapstructure: "PORT"`
	GRPCServerAddress      string        `mapstructure: "GRPCSERVERADDRESS"`
	AccessTokenPrivateKey  string        `mapstructure: "ACCESSTOKENPRIVATEKEY"`
	AccessTokenPublicKey   string        `mapstructure: "ACCESSTOKENPUBLICKEY"`
	RefreshTokenPrivateKey string        `mapstructure: "REFRESHTOKENPRIVATEKEY"`
	RefreshTokenPublicKey  string        `mapstructure: "REFRESHTOKENPUBLICKEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure: "ACCESSTOKENEXPIRESIN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure: "REFRESHTOKENEXPIRESIN"`
	AccessTokenMaxAge      int           `mapstructure: "ACCESSTOKENMAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure: "REFRESHTOKENMAXAGE"`
	Origin                 string        `mapstructure: "ORIGIN"`
	EmailFrom              string        `mapstructure: "EMAILFROM"`
	SMTPHost               string        `mapstructure: "SMTPHOST"`
	SMTPPass               string        `mapstructure: "SMTPPASS"`
	SMTPPort               int           `mapstructure: "SMTPPORT"`
	SMTPUser               string        `mapstructure: "SMTPUSER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
