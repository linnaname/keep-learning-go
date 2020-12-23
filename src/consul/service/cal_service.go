package service

import (
	"context"
	"keep-learning-go/src/consul/api"
	"math/rand"
	"time"
)

type CalService struct {
}

func (s *CalService) Add(ctx context.Context, request *api.Request) (*api.Response, error) {
	// 50% 概率 sleep，模拟超时场景
	if rand.Int()%2 == 0 {
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	response := &api.Response{
		V: request.A + request.B,
	}
	return response, nil
}
