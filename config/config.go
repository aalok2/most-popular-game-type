package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    RedisAddress string
    MongoURI     string
    MongoDBName string
}

func LoadConfig() *Config {
    viper.SetConfigFile(".env")
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    return &Config{
        RedisAddress: viper.GetString("REDIS_ADDRESS"),
        MongoURI:     viper.GetString("MONGO_URI"),
	    MongoDBName:  viper.GetString("MONGO_DB_NAME"),
    }
}
