package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisData interface {
	MarshalBinary() ([]byte, error)
}

func CreateRedisClient(connectionString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		fmt.Println("problem parsing redis connection string:", err)
		return nil, err
	}

	rdb := redis.NewClient(opt)
	return rdb, nil
}

func GetAndSetValue[T any](ctx context.Context, client *redis.Client, key string, notExistsFunc func() (T, error)) (T, error) {
	tctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	val, err := client.Get(tctx, key).Result()
	var result T
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
		case errors.Is(err, redis.Nil):
			value, err := notExistsFunc()
			if err != nil {
				return result, err
			}

			js, err := encodedToJSONString(value)
			if err != nil {
				return result, err
			}

			res, err := client.Set(ctx, key, js, 24*time.Hour).Result()
			if err != nil {
				return result, err
			}
			fmt.Println(res)
			return value, nil
		default:
			return result, err
		}

	}

	err = decodedToType(val, &result)
	if err != nil {
		return result, err
	}
	fmt.Println("Found val from redis")

	return result, nil
}

func decodedToType[T any](val string, dst *T) error {
	err := json.Unmarshal([]byte(val), dst)
	if err != nil {
		return err
	}

	return nil
}

func encodedToJSONString[T any](value T) (string, error) {
	js, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(js), nil
}
