package main

import (
	"fmt"

	"github.com/MdSohelMia/go-cache/cache"
)

func main() {
	CacheTest()
}

type Cache struct {
	Id    string
	Value interface{}
}

func CacheTest() {

	redis := cache.NewRedis(cache.RedisConfig{
		Host:     "127.0.0.1",
		Password: "",
		Database: 0,
		Port:     6379,
	})
	c := Cache{Id: "3140", Value: "hello"}

	cache := cache.NewCache(redis)

	val := cache.RememberForever("videos", func() any {
		return c
	})

	fmt.Println(val)
}
