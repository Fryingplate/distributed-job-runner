package worker

import (
	"log"
	"time"

	"distributed-job-runner/internal/job"
	"distributed-job-runner/internal/queue"
)

type Worker struct {
	Service *job.Service
	Queue   queue.Queue
}

func (w *Worker) Start() {

	for {

		jobID, err := w.Queue.Consume()

		if err != nil {
			log.Println("Queue consume error:", err)
			continue
		}

		w.Process(jobID)
	}
}

func (w *Worker) Process(
	jobID string,
) {

	j, err := w.Service.GetByID(jobID)

	if err != nil {
		log.Println("Get job error:", err)
		return
	}

	err = w.Service.UpdateStatus(
		jobID,
		"RUNNING",
	)

	if err != nil {
		log.Println("Update RUNNING error:", err)
		return
	}

	log.Println(
		"Processing:",
		j.Name,
	)

	// Simulate actual work
	time.Sleep(5 * time.Second)

	err = w.Service.UpdateStatus(
		jobID,
		"COMPLETED",
	)

	if err != nil {
		log.Println("Update COMPLETED error:", err)
		return
	}

	log.Println(
		"Completed:",
		j.Name,
	)
}
