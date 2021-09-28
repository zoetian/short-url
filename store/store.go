package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("Successfully start redis: pong msg = {%s}", pong)
	storeService.redisClient = redisClient

	return storeService
}

func SaveUrlMapping(shortUrl, longUrl, userID string) {
	err := storeService.redisClient.Set(shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to save key url.\nError: %v \nShortUrl: %s \nOriginalUrl: %s\n", err, shortUrl, longUrl))
	}
}

func FetchOriginalUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch original url\n Error: %v \nShortUrl: %s\n", err, shortUrl))
	}
	return result
}
