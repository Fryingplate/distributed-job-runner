package job

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) create(j Job) error {
	_, err := r.DB.Exec(

		`
		INSERT INTO jobs
		(
			id,
			name,
			status,
			created_at
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4
		)
		`,
		j.ID,
		j.Name,
		j.Status,
		j.CreatedAt,
	)

	return err
}

func (r *Repository) GetByID(
	id string,
) (Job, error) {

	var j Job

	err := r.DB.QueryRow(
		`
		SELECT
			id,
			name,
			status,
			created_at
		FROM jobs
		WHERE id = $1
		`,
		id,
	).Scan(
		&j.ID,
		&j.Name,
		&j.Status,
		&j.CreatedAt,
	)

	return j, err
}

func (r *Repository) UpdateStatus(
	id string,
	status string,
) error {

	_, err := r.DB.Exec(
		`
		UPDATE jobs
		SET status = $1
		WHERE id = $2
		`,
		status,
		id,
	)

	return err
}
