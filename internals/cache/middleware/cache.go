package cache

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"
	"tools/internals/models"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Connection *redis.Client
}

func NewCache() *Cache {
	connection := NewRedisConnection()
	return &Cache{Connection: connection}
}

func NewRedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("%v", ErrorPinging)
		return nil
	}
	log.Printf("Successfully started a new redis connection: %s\n", ping)
	return client
}

func (c *Cache) FindBanner(key string) (*models.Banner, error) {
	var val []byte
	banner := &models.Banner{}
	err := c.Connection.Get(context.Background(), key).Scan(&val)
	if err == redis.Nil {
		log.Printf("%v", ErrorNoBannerFound)
		err = ErrorNilReturned
		return nil, err
	}
	log.Println("Took the value from cache")
	banner.Content = json.RawMessage(val)
	return banner, nil
}

func (c *Cache) AddFromRepo(input *models.Banner, request *models.BannerGetRequest) error {
	newKey := strconv.FormatUint(uint64(request.TagID), 10) + strconv.FormatUint(uint64(request.FeatureID), 10)
	data, err := input.Content.MarshalJSON()
	if err != nil {
		log.Printf("%v", ErrorMarshalProblem)
		return err
	}
	err = c.Connection.Set(context.Background(), newKey, data, 5*time.Minute).Err()
	if err != nil {
		log.Printf("%v", ErrorSetRedis)
		return err
	}
	log.Println("Added the value to the cache")
	return nil
}
