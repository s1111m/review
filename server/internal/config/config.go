package config

import (
	"fmt"
	"runtime"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cfg *Config = &Config{}

type Config struct {
	PROTO         string `mapstructure:"GRPC_SERVER_BIND_PROTOCOL"`
	PORT          string `mapstructure:"GRPC_SERVER_BIND_PORT"`
	LOGGING_LEVEL string `mapstructure:"LOGGING_LEVEL"`
	MAX_THREADS   int    `mapstructure:"MAX_THREADS"`
}

func (c *Config) SetLoggingLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {

	}
	logrus.SetLevel(lvl)
}

func (c *Config) GetLoggingLevel() logrus.Level {
	return logrus.GetLevel()
}

func GetMaxThreads() int {
	return Cfg.MAX_THREADS
}

func init() {
	viper.AddConfigPath(".")

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("PROTO", "tcp")
	viper.SetDefault("PORT", ":50051")
	viper.SetDefault("LOGGING_LEVEL", "error")
	fmt.Println(runtime.NumCPU())
	viper.SetDefault("MAX_THREADS", runtime.NumCPU())

	err := viper.ReadInConfig()
	if err != nil {
		panic("Cannot Read Config")
	}
	err = viper.Unmarshal(&Cfg)
	runtime.GOMAXPROCS(Cfg.MAX_THREADS)
	Cfg.SetLoggingLevel(Cfg.LOGGING_LEVEL)
}
