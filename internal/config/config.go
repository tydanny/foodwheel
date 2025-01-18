package config

import (
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var once sync.Once

func InitializeConfig(log *zap.Logger) {
	once.Do(func() {
		log := zap.L().Named("config")

		// Initialize the configuration
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		// Set default values
		viper.SetDefault("pg.dbaddr", "localhost:5432")
		viper.SetDefault("pg.dbuser", "postgres")
		viper.SetDefault("pg.dbpass", "postgres")

		// Read the configuration
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Info("config file not found; using defaults")
			} else {
				log.Fatal("failed to read config file", zap.Error(err))
			}
		}
	})
}
