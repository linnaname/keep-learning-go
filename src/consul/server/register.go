package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
	"strconv"
	"time"
)

type ConsulRegister struct {
	Address                        string
	Service                        string
	Tag                            []string
	Port                           int
	BalanceFactor                  int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

func NewConsulRegister() *ConsulRegister {
	return &ConsulRegister{
		Address:                        "127.0.0.1:8500",
		Service:                        "unknown",
		Tag:                            []string{},
		Port:                           3000,
		BalanceFactor:                  100,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
}

func (r *ConsulRegister) Register() error {
	config := api.DefaultConfig()
	config.Address = r.Address
	client, err := api.NewClient(config)
	if err != nil {
		log.Println("consul client error : ", err)
		return err
	}

	ip := localIP()
	//元数据
	registration := new(api.AgentServiceRegistration)
	registration.ID = fmt.Sprintf("%v-%v-%v", r.Service, ip, r.Port) // 服务节点的名称
	registration.Name = fmt.Sprintf("%v", r.Service)                 // 服务名称
	registration.Port = r.Port                                       // 服务端口
	registration.Tags = r.Tag                                        // tag，可以为空
	registration.Address = ip                                        // 服务 IP
	registration.Meta = map[string]string{
		"balanceFactor": strconv.Itoa(r.BalanceFactor),
	}

	// 健康检查
	registration.Check = &api.AgentServiceCheck{
		Interval:                       r.Interval.String(),                            // 健康检查间隔
		DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),      //check失败后30秒删除本服务，注销时间，相当于过期时间
		GRPC:                           fmt.Sprintf("%v:%v/%v", ip, r.Port, r.Service), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	//注册
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Println("register server error : ", err)
		return err
	}
	return nil
}

/**
本地ip
*/
func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
