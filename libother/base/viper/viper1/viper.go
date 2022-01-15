package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func main() {

}

type config struct {
	v *viper.Viper
}

func Do() {
	c := &config{}
	WatchConfig(c)
}

func LoadConfigFromYaml(c *config) error {
	c.v = viper.New()
	//设置配置文件的名字
	c.v.SetConfigName("config")

	c.v.AddConfigPath("/Users/xulei/jungle/golangworkspace/helen/vipers")

	//设置配置文件类型
	c.v.SetConfigType("yaml")

	if err := c.v.ReadInConfig(); err != nil {
		fmt.Print(err)
		return err
	}

	log.Printf("age: %s, name: %s \n", c.v.Get("information.age"), c.v.Get("information.name"))
	return nil
}

//监听配置文件的修改和变动
func WatchConfig(c *config) error {
	if err := LoadConfigFromYaml(c); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	c.v.WatchConfig()

	//监听回调函数
	watch := func(e fsnotify.Event) {
		log.Printf("Config file is changed: %s \n", e.String())
		log.Printf("age: %s, name: %s \n", c.v.Get("information.age"), c.v.Get("information.name"))
		cancel()
	}

	c.v.OnConfigChange(watch)
	<-ctx.Done()
	return nil
}
