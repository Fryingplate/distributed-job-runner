package main

import (
	"log"
	"net/http"

	"distributed-job-runner/internal/database"
	"distributed-job-runner/internal/job"
	"distributed-job-runner/internal/queue"

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

	handler := &job.Handler{
		Service: service,
		Queue:   redisQueue,
	}

	http.HandleFunc(
		"/jobs",
		handler.CreateJob,
	)

	log.Println(
		"API running on :8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			nil,
		),
	)
}
