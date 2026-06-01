package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	Client *redis.Client
}

func (r *RedisQueue) Publish(jobID string) error {

	return r.Client.RPush(
		context.Background(),
		"jobs",
		jobID,
	).Err()
}

func (r *RedisQueue) Consume() (string, error) {

	result, err := r.Client.BLPop(
		context.Background(),
		0,
		"jobs",
	).Result()

	if err != nil {
		return "", err
	}

	return result[1], nil
}
