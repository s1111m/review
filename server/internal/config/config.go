package config

import (
	"runtime"
	"runtime/debug"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cfg *Config = &Config{}
var Logger *logrus.Entry

type Config struct {
	PROTO         string `mapstructure:"GRPC_SERVER_BIND_PROTOCOL"`
	PORT          string `mapstructure:"GRPC_SERVER_BIND_PORT"`
	LOGGING_LEVEL string `mapstructure:"LOGGING_LEVEL"`
	MAX_THREADS   int    `mapstructure:"MAX_THREADS"` // максимальное число горутин за раз, ограничиваемся архитектурой проца, либо выставляем вручную
}

// враппер для установки уровня логирования
func (c *Config) SetLoggingLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {

	}
	logrus.SetLevel(lvl)
}

// геттер на уровень логирования
func (c *Config) GetLoggingLevel() logrus.Level {
	return logrus.GetLevel()
}

//геттер на число потоков
func GetMaxThreads() int {
	return Cfg.MAX_THREADS
}

func init() {
	// вычитали конфиг. выставили дефолты
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("PROTO", "tcp")
	viper.SetDefault("PORT", ":50051")
	viper.SetDefault("LOGGING_LEVEL", "error")
	viper.SetDefault("MAX_THREADS", runtime.NumCPU())

	err := viper.ReadInConfig()
	if err != nil {
		panic("Cannot Read Config")
	}
	err = viper.Unmarshal(&Cfg)
	runtime.GOMAXPROCS(Cfg.MAX_THREADS)
	Cfg.SetLoggingLevel(Cfg.LOGGING_LEVEL)
	//set stacktrace to default
	Logger = logrus.WithField("stack", string(debug.Stack()))

	//Logger.Formatter = NewGelf("server")
	//hook := graylog.NewGraylogHook("localhost:12201", map[string]interface{}{})
	//Logger.AddHook(hook)

}
