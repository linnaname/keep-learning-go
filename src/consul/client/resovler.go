package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"net"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	regexConsul, _ = regexp.Compile("^([A-z0-9.]+)(:[0-9]{1,5})?/([A-z_]+)$")
)

type ConsulResolver struct {
	address              string
	wg                   sync.WaitGroup
	cc                   resolver.ClientConn
	name                 string
	disableServiceConfig bool
	Ch                   chan int
}

func NewConsulResolver() resolver.Builder {
	return &ConsulResolver{}
}

func (cb *ConsulResolver) ResolveNow(options resolver.ResolveNowOptions) {
	cb.Ch <- 1
}

func (cb *ConsulResolver) Close() {
}

func (cb *ConsulResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	host, port, name, err := parseTarget(fmt.Sprintf("%s/%s", target.Authority, target.Endpoint))
	if err != nil {
		fmt.Println("parse err")
		return nil, err
	}
	fmt.Println(fmt.Sprintf("consul service ==> host:%s, port%s, name:%s", host, port, name))

	cr := &ConsulResolver{
		address:              fmt.Sprintf("%s%s", host, port),
		name:                 name,
		cc:                   cc,
		disableServiceConfig: opts.DisableServiceConfig,
		Ch:                   make(chan int, 0),
	}
	go cr.watcher()
	return cr, nil
}

func (cb *ConsulResolver) watcher() {
	fmt.Printf("calling [%s] consul watcher\n", cb.name)
	config := api.DefaultConfig()
	config.Address = cb.address
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Printf("error create consul client: %v\n", err)
		return
	}
	t := time.NewTicker(2000 * time.Millisecond)
	defer func() {
		fmt.Println("defer done")
	}()
	for {
		select {
		case <-t.C:
			//fmt.Println("定时")
		case <-cb.Ch:
			//fmt.Println("ch call")
		}
		//api添加了 lastIndex   consul api中并不兼容附带lastIndex的查询
		services, _, err := client.Health().Service(cb.name, "", true, &api.QueryOptions{})
		if err != nil {
			fmt.Printf("error retrieving instances from Consul: %v", err)
		}

		newAddrs := make([]resolver.Address, 0)
		for _, service := range services {
			addr := net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))
			newAddrs = append(newAddrs, resolver.Address{
				Addr: addr,
				//type：不能是grpclb，grpclb在处理链接时会删除最后一个链接地址，不用设置即可 详见=> balancer_conn_wrappers => updateClientConnState
				ServerName: service.Service.Service,
			})
		}
		//cb.cc.NewAddress(newAddrs)
		//cb.cc.NewServiceConfig(cr.name)
		cb.cc.UpdateState(resolver.State{Addresses: newAddrs})
	}

}

func (cb *ConsulResolver) Scheme() string {
	return "consul grpc"
}

func parseTarget(target string) (host, port, name string, err error) {
	if target == "" {
		return "", "", "", errors.New("illegal argument")
	}

	if !regexConsul.MatchString(target) {
		return "", "", "", errors.New("no match")
	}

	groups := regexConsul.FindStringSubmatch(target)
	host = groups[1]
	port = groups[2]
	name = groups[3]
	return host, port, name, nil
}
