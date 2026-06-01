package main

import (
	"log"

	"distributed-job-runner/internal/database"
	"distributed-job-runner/internal/job"
	"distributed-job-runner/internal/queue"
	"distributed-job-runner/internal/worker"

	"github.com/redis/go-redis/v9"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repo := &job.Repository{
		DB: db,
	}

	service := &job.Service{
		Repo: repo,
	}

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)

	redisQueue := &queue.RedisQueue{
		Client: redisClient,
	}

	w := &worker.Worker{
		Service: service,
		Queue:   redisQueue,
	}

	log.Println(
		"Worker started",
	)

	w.Start()
}
