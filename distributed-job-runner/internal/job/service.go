package job

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Repo *Repository
}

func (s *Service) Create(
	name string,
) (Job, error) {

	j := Job{
		ID:        uuid.NewString(),
		Name:      name,
		Status:    "PENDING",
		CreatedAt: time.Now(),
	}

	err := s.Repo.create(j)

	if err != nil {
		return Job{}, err
	}

	return j, nil
}

func (s *Service) GetByID(
	id string,
) (Job, error) {

	return s.Repo.GetByID(id)
}

func (s *Service) UpdateStatus(
	id string,
	status string,
) error {

	return s.Repo.UpdateStatus(
		id,
		status,
	)
}
