package repositories

import (
	"database/sql"
	m "prem/internal/models"
)

// CREATE
func CreateJob(db *sql.DB, job m.Job) error {

	query := `INSERT INTO jobs (title, description, apply_url) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, job.Title, job.Description, job.ApplyURL)
	return err
}

// READ (all)
func GetJobs(db *sql.DB) ([]m.Job, error) {
	rows, err := db.Query(`SELECT id, title, description, apply_url FROM jobs`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []m.Job
	for rows.Next() {
		var job m.Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.ApplyURL)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

// UPDATE
func UpdateJob(db *sql.DB, job m.Job) error {
	query := `UPDATE jobs SET title=$1, description=$2, apply_url=$3 WHERE id=$4`
	_, err := db.Exec(query, job.Title, job.Description, job.ApplyURL, job.ID)
	return err
}

// DELETE
func DeleteJob(db *sql.DB, id int) error {
	query := `DELETE FROM jobs WHERE id=$1`
	_, err := db.Exec(query, id)
	return err
}
