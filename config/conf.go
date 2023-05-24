package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type App struct {
	PageSize int    `json:"page_size"`
	Location string `json:"loacation"`
}

type Server struct {
	HTTPPort     int           `json:"http_port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type Mysql struct {
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	DbName      string `json:"db_name"`
	TablePrefix string `json:"table_prefix"`
}

type Redis struct {
	RedisHost  string `json:"redis_host"`
	RedisIndex string `json:"redis_index"`
}

// Config 服务端配置数据结构
type Config struct {
	RunMode string `json:"run_mode"`
	APP     App    `json:"app"`
	Server  Server `json:"server"`
	MySQL   Mysql  `json:"my_sql"`
	Redis   Redis  `json:"redis"`
}

var ServerConfig = Config{}

func InitConfig(confPath string) {
	viper.SetConfigFile(confPath)
	if err := viper.ReadInConfig(); err != nil {
		// 处理读取配置文件失败的情况
		log.Panicf("read conf error %s", err.Error())
	}

	if err := viper.Unmarshal(&ServerConfig); err != nil {
		log.Panicf("unmarshal conf error %s", err.Error())
	}
	log.Printf("conf %+v", ServerConfig)
}
