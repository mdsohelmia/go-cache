package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type RedisConfig struct {
	Host     string
	Password string
	Port     int
	Database int
}

type Redis struct {
	config RedisConfig
	client *redis.Client
}

// Redis Config
func NewRedis(config RedisConfig) *Redis {
	redis := &Redis{
		config: config,
	}
	redis.connect()
	return redis
}

func (r *Redis) connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.config.Host, r.config.Port), //Configure
		Password: r.config.Password,                                  // no password set
		DB:       r.config.Database,                                  // use default DB
	})

	r.client = rdb
}

func (redis Redis) get(key string) (any, error) {
	return nil, nil
}
func (redis Redis) has(key string) bool {
	return false
}
func (redis Redis) remember(key string, ttl time.Time, callback CallbackFunc) any {
	return nil
}
func (redis Redis) rememberForever(key string, callback CallbackFunc) any {
	//Callback return any data
	value := callback()

	err := redis.client.Set(ctx, key, value, time.Minute*0).Err()

	if err != nil {
		fmt.Println(err)
		return nil
	}
	val, err := redis.client.Get(ctx, key).Result()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return val
}
func (redis Redis) pull(key string) any {
	return nil
}
func (redis Redis) put(key string, value any, ttl ...time.Time) {

}
func (redis Redis) add(key string, value any, ttl ...time.Time) {

}
func (redis Redis) forever(key string, value any) {

}
func (redis Redis) forget(key string) error {
	return nil
}
func (redis Redis) flush() string {
	return "flush"
}
