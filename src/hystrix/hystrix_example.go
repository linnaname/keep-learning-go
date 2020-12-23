package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	metricCollector "github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/afex/hystrix-go/plugins"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	//限流熔断配置
	hystrix.ConfigureCommand("test", hystrix.CommandConfig{
		//超时时间设置  单位毫秒
		Timeout: 1000,
		//最大请求数
		MaxConcurrentRequests: 100,
		//错误率
		ErrorPercentThreshold: 25,
		//请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
		RequestVolumeThreshold: 5,
		//过多长时间，熔断器再次检测是否开启。单位毫秒
		SleepWindow: 1,
	})

	//dashboard
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	//metric collector
	c, err := plugins.InitializeStatsdCollector(&plugins.StatsdCollectorConfig{
		StatsdAddr: "localhost:8125",
		Prefix:     "myapp.hystrix",
	})
	if err != nil {
		log.Fatalf("could not initialize statsd client: %v", err)
	}
	metricCollector.Registry.Register(c.NewStatsdCollector)

	hystrix.Go("test", func() error {
		//把网断了就可以看到效果了
		_, err := http.Get("https://www.baidu.com/")
		if err != nil {
			fmt.Println("get error")
			return err
		}
		return nil
	}, func(err error) error {
		fmt.Println("get an error, handle it")
		return nil
	})

	time.Sleep(2 * time.Second)

}
