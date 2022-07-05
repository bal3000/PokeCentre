package data

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
)

func CreateRedisClient(connectionString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		fmt.Println("problem parsing redis connection string:", err)
		return nil, err
	}

	rdb := redis.NewClient(opt)
	return rdb, nil
}

func GetOrSetValue[T any](ctx context.Context, client *redis.Client, key string, notExistsFunc func() (T, error)) (T, error) {
	tctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	val, err := client.Get(tctx, key).Result()
	var result T
	fmt.Println("val from redis:", val)
	if err != nil {
		if err == redis.Nil {
			value, err := notExistsFunc()
			if err != nil {
				return result, err
			}
			res, err := client.Set(ctx, key, value, 24*time.Hour).Result()
			if err != nil {
				return result, err
			}
			fmt.Println(res)
			return value, nil
		}

		return result, err
	}

	err = convertRedisValToType(val, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func convertRedisValToType[T any](val string, dst T) error {
	r := strings.NewReader(val)
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)

	if err != nil {
		return err
	}

	return nil
}
