package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)
import "github.com/go-redis/redis/v8"

var ctx = context.Background()

func TestSetAndGet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	assert.NoError(t, err)

	val, err := rdb.Get(ctx, "key").Result()
	assert.NoError(t, err)
	assert.NotEmpty(t, val)
	assert.Equal(t, val, "value")

	val2, err := rdb.Get(ctx, "key2").Result()
	assert.Equal(t, err, redis.Nil)
	assert.Empty(t, val2)
}

func TestDatastruce(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	set, err := rdb.SetNX(ctx, "lin", "nana", 10*time.Second).Result()
	assert.NoError(t, err)
	assert.True(t, set)

	vals, err := rdb.Sort(ctx, "mylist", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
	assert.NoError(t, err)
	assert.Len(t, vals, 2)
	assert.Contains(t, vals, "1")

	res, err := rdb.Do(ctx, "set", "USA", "CN").Result()
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
