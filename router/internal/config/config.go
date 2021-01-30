package config

import (
	"runtime/debug"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cfg *Config = &Config{}
var Logger *logrus.Entry

type Config struct {
	BIND_PORT         int    `mapstructure:"BIND_PORT"`
	DB_PATH           string `mapstructure:"DB_PATH"`
	GRPC_SERVICE_ADDR string `mapstructure:"GRPC_SERVICE_ADDR"`
	GRPC_SERVICE_PORT string `mapstructure:"GRPC_SERVICE_PORT"`
	LOGGING_LEVEL     string `mapstructure:"LOGGING_LEVEL"`
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

func init() {
	viper.AddConfigPath(".")

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("BIND_PORT", 4444)
	viper.SetDefault("LOGGING_LEVEL", "error")
	viper.SetDefault("DB_PATH", "./gorm.db")
	viper.SetDefault("GRPC_SERVICE_ADDR", "server")
	viper.SetDefault("GRPC_SERVICE_PORT", "50051")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Cannot Read Config")
	}
	err = viper.Unmarshal(&Cfg)
	Cfg.SetLoggingLevel(Cfg.LOGGING_LEVEL)
	//set stacktrace to default
	Logger = logrus.WithField("stack", string(debug.Stack()))
	//Logger.SetFormatter(formatters.NewGelf("router"))
	//hook := graylog.NewGraylogHook("localhost:12201", map[string]interface{}{})
	//Logger.AddHook(hook)

}
