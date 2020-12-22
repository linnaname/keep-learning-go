package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

func init() {
	// 绑定环境变量
	viper.AutomaticEnv()
}

type Config struct {
	AppName  string
	LogLevel string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

func main() {
	viper.SetConfigName("test")
	viper.SetConfigType("toml") // REQUIRED if the config file does not have the extension in the name
	//可以设置多个，viper 会根据设置顺序依次查找；
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))

	fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports: ", viper.GetIntSlice("server.ports"))
	fmt.Println("timeout: ", viper.GetDuration("server.timeout"))

	if viper.IsSet("redis.port") {
		fmt.Println("redis.port is set")
	} else {
		fmt.Println("redis.port is not set")
	}

	fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	fmt.Println("all settings: ", viper.AllSettings())

	fmt.Println("GOPATH: ", viper.Get("GOPATH"))

	var c Config
	viper.Unmarshal(&c)
	fmt.Println(c.MySQL)

	viper.Set("mysql.ext", "ext")
	err = viper.WriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})
	fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	time.Sleep(time.Second * 30)
	fmt.Println("redis port after sleep: ", viper.Get("redis.port"))

}
