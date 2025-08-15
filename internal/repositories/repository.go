package repository

import (
	"database/sql"
	"log"

	"github.com/chetan-sharma-exe/sarkari-backend/internal/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetJobs() []models.Job {
	rows, err := r.DB.Query(`SELECT id, title, description, apply_url, last_date, date_posted FROM jobs`)
	if err != nil {
		log.Println("GetJobs error:", err)
		return nil
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var j models.Job
		rows.Scan(&j.ID, &j.Title, &j.Description, &j.Url, &j.LastDate, &j.DatePosted)
		jobs = append(jobs, j)
	}
	return jobs
}

func (r *Repository) GetResults() []models.Result {
	rows, err := r.DB.Query(`SELECT id, title, result_link, exam_date, date_posted FROM results`)
	if err != nil {
		log.Println("GetResults error:", err)
		return nil
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var res models.Result
		rows.Scan(&res.ID, &res.Title, &res.Url, &res.ExamDate, &res.DatePosted)
		results = append(results, res)
	}
	return results
}

func (r *Repository) GetAdmitCards() []models.AdmitCard {
	rows, err := r.DB.Query(`SELECT id, title, download_link, exam_date, date_posted FROM admit_cards`)
	if err != nil {
		log.Println("GetAdmitCards error:", err)
		return nil
	}
	defer rows.Close()

	var admitCards []models.AdmitCard
	for rows.Next() {
		var a models.AdmitCard
		rows.Scan(&a.ID, &a.Title, &a.Url, &a.ExamDate, &a.DatePosted)
		admitCards = append(admitCards, a)
	}
	return admitCards
}

func (r *Repository) GetVImpLinks() []models.VImpLink {
	rows, err := r.DB.Query(`SELECT id, title, link FROM vimps_links`)
	if err != nil {
		log.Println("GetVimpLinks error:", err)
		return nil
	}
	defer rows.Close()

	var links []models.VImpLink
	for rows.Next() {
		var v models.VImpLink
		rows.Scan(&v.ID, &v.Title, &v.Url)
		links = append(links, v)
	}
	return links
}

func (r *Repository) GetImpLinks() []models.ImpLink {
	rows, err := r.DB.Query(`SELECT id, title, link FROM imp_links`)
	if err != nil {
		log.Println("GetImpLinks error:", err)
		return nil
	}
	defer rows.Close()

	var links []models.ImpLink
	for rows.Next() {
		var i models.ImpLink
		rows.Scan(&i.ID, &i.Title, &i.Url)
		links = append(links, i)
	}
	return links
}

func (r *Repository) GetAllData() models.AllData {
	return models.AllData{
		Jobs:       r.GetJobs(),
		Results:    r.GetResults(),
		AdmitCards: r.GetAdmitCards(),
		VimpLinks:  r.GetVImpLinks(),
		ImpLinks:   r.GetImpLinks(),
	}
}

func (r *Repository) InsertJob(job models.Job) error {
	_, err := r.DB.Exec(`
		INSERT INTO jobs (title, description, apply_url, last_date) 
		VALUES ($1, $2, $3, $4)`,
		job.Title, job.Description, job.Url, job.LastDate,
	)
	return err
}

func (r *Repository) InsertResult(res models.Result) error {
	_, err := r.DB.Exec(`
		INSERT INTO results (title, result_link, exam_date) 
		VALUES ($1, $2, $3)`,
		res.Title, res.Url, res.ExamDate,
	)
	return err
}

func (r *Repository) InsertAdmitCard(admit models.AdmitCard) error {
	_, err := r.DB.Exec(`
		INSERT INTO admit_cards (title, download_link, exam_date) 
		VALUES ($1, $2, $3)`,
		admit.Title, admit.Url, admit.ExamDate,
	)
	return err
}

func (r *Repository) InsertVimpLink(link models.VImpLink) error {
	_, err := r.DB.Exec(`
		INSERT INTO vimps_links (title, link) 
		VALUES ($1, $2)`,
		link.Title, link.Url,
	)
	return err
}

func (r *Repository) InsertImpLink(link models.ImpLink) error {
	_, err := r.DB.Exec(`
		INSERT INTO imp_links (title, link) 
		VALUES ($1, $2)`,
		link.Title, link.Url,
	)
	return err
}

func (r *Repository) InsertAllData(data models.AllData) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Step 1: Clear all tables (TRUNCATE resets serial IDs)
	_, err = tx.Exec(`TRUNCATE jobs, results, admit_cards, vimps_links, imp_links RESTART IDENTITY CASCADE`)
	if err != nil {
		return err
	}

	// Insert jobs
	for _, job := range data.Jobs {
		_, err := tx.Exec(`
			INSERT INTO jobs (title, description, apply_url, last_date) 
			VALUES ($1, $2, $3, $4)`,
			job.Title, job.Description, job.Url, job.LastDate,
		)
		if err != nil {
			return err
		}
	}

	// Insert results
	for _, res := range data.Results {
		_, err := tx.Exec(`
			INSERT INTO results (title, result_link, exam_date) 
			VALUES ($1, $2, $3)`,
			res.Title, res.Url, res.ExamDate,
		)
		if err != nil {
			return err
		}
	}

	// Insert admit cards
	for _, admit := range data.AdmitCards {
		_, err := tx.Exec(`
			INSERT INTO admit_cards (title, download_link, exam_date) 
			VALUES ($1, $2, $3)`,
			admit.Title, admit.Url, admit.ExamDate,
		)
		if err != nil {
			return err
		}
	}

	// Insert vimps_links
	for _, vlink := range data.VimpLinks {
		_, err := tx.Exec(`
			INSERT INTO vimps_links (title, link) 
			VALUES ($1, $2)`,
			vlink.Title, vlink.Url,
		)
		if err != nil {
			return err
		}
	}

	// Insert imp_links
	for _, ilink := range data.ImpLinks {
		_, err := tx.Exec(`
			INSERT INTO imp_links (title, link) 
			VALUES ($1, $2)`,
			ilink.Title, ilink.Url,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
