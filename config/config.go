package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongoDBURI        string `mapstructure:"MONGODB_URI"`
	RedisURI          string `mapstructure:"REDIS_URI"`
	Port              string `mapstructure:"PORT"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	Origin            string `mapstructure:"ORIGIN"`
	BucketRegion      string `mapstructure:"BUCKET_REGION"`
	BucketEndpointURL string `mapstructure:"BUCKET_ENDPOINT_URL"`
	BucketAccessKey   string `mapstructure:"BUCKET_ACCESS_KEY"`
	BucketAccessID    string `mapstructure:"BUCKET_ACCESS_ID"`
	BucketName        string `mapstructure:"BUCKET_NAME"`
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
