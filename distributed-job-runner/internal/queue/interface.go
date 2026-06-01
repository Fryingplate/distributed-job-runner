package queue

type Queue interface {
	Publish(jobID string) error
	Consume() (string, error)
}
