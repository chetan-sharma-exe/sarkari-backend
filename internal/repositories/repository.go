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
	rows, err := r.DB.Query(`SELECT id, title, description, url, last_date, date_posted FROM jobs`)
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
	rows, err := r.DB.Query(`SELECT id, title, url, exam_date, date_posted FROM results`)
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
	rows, err := r.DB.Query(`SELECT id, title, url, exam_date, date_posted FROM admit_cards`)
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
	rows, err := r.DB.Query(`SELECT id, title, url FROM vimps_links`)
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
	rows, err := r.DB.Query(`SELECT id, title, url FROM imp_links`)
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

func (r *Repository) GetNaukriForms() []models.NaukriForm {
	rows, err := r.DB.Query(`SELECT id, title, url FROM naukri_forms`)
	if err != nil {
		log.Println("GetNaukriForms error:", err)
		return nil
	}
	defer rows.Close()

	var forms []models.NaukriForm
	for rows.Next() {
		var f models.NaukriForm
		rows.Scan(&f.ID, &f.Title, &f.Url)
		forms = append(forms, f)
	}
	return forms
}

func (r *Repository) GetAdmissions() []models.Admission {
	rows, err := r.DB.Query(`SELECT id, title, url FROM admissions`)
	if err != nil {
		log.Println("GetAdmissions error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.Admission
	for rows.Next() {
		var a models.Admission
		rows.Scan(&a.ID, &a.Title, &a.Url)
		list = append(list, a)
	}
	return list
}

func (r *Repository) GetRegularForms() []models.RegularForm {
	rows, err := r.DB.Query(`SELECT id, title, url FROM regular_forms`)
	if err != nil {
		log.Println("GetRegularForms error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.RegularForm
	for rows.Next() {
		var f models.RegularForm
		rows.Scan(&f.ID, &f.Title, &f.Url)
		list = append(list, f)
	}
	return list
}

func (r *Repository) GetOfflineForms() []models.OfflineForm {
	rows, err := r.DB.Query(`SELECT id, title, url FROM offline_forms`)
	if err != nil {
		log.Println("GetOfflineForms error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.OfflineForm
	for rows.Next() {
		var f models.OfflineForm
		rows.Scan(&f.ID, &f.Title, &f.Url)
		list = append(list, f)
	}
	return list
}

func (r *Repository) GetAnswerKeys() []models.AnswerKey {
	rows, err := r.DB.Query(`SELECT id, title, url FROM answer_keys`)
	if err != nil {
		log.Println("GetAnswerKeys error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.AnswerKey
	for rows.Next() {
		var a models.AnswerKey
		rows.Scan(&a.ID, &a.Title, &a.Url)
		list = append(list, a)
	}
	return list
}

func (r *Repository) GetSyllabus() []models.Syllabus {
	rows, err := r.DB.Query(`SELECT id, title, url FROM syllabus`)
	if err != nil {
		log.Println("GetSyllabus error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.Syllabus
	for rows.Next() {
		var s models.Syllabus
		rows.Scan(&s.ID, &s.Title, &s.Url)
		list = append(list, s)
	}
	return list
}

func (r *Repository) GetSarkariYojna() []models.SarkariYojna {
	rows, err := r.DB.Query(`SELECT id, title, url FROM sarkari_yojna`)
	if err != nil {
		log.Println("GetSarkariYojna error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.SarkariYojna
	for rows.Next() {
		var s models.SarkariYojna
		rows.Scan(&s.ID, &s.Title, &s.Url)
		list = append(list, s)
	}
	return list
}

func (r *Repository) GetVerification() []models.Verification {
	rows, err := r.DB.Query(`SELECT id, title, url FROM verification`)
	if err != nil {
		log.Println("GetVerification error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.Verification
	for rows.Next() {
		var v models.Verification
		rows.Scan(&v.ID, &v.Title, &v.Url)
		list = append(list, v)
	}
	return list
}

func (r *Repository) GetUpcoming() []models.Upcoming {
	rows, err := r.DB.Query(`SELECT id, title, url FROM upcoming`)
	if err != nil {
		log.Println("GetUpcoming error:", err)
		return nil
	}
	defer rows.Close()

	var list []models.Upcoming
	for rows.Next() {
		var u models.Upcoming
		rows.Scan(&u.ID, &u.Title, &u.Url)
		list = append(list, u)
	}
	return list
}

func (r *Repository) GetAllData() models.AllData {
	return models.AllData{
		Jobs:             r.GetJobs(),
		Results:          r.GetResults(),
		AdmitCards:       r.GetAdmitCards(),
		VimpLinks:        r.GetVImpLinks(),
		ImpLinks:         r.GetImpLinks(),
		NaukriForms:      r.GetNaukriForms(),
		Admissions:       r.GetAdmissions(),
		RegularForms:     r.GetRegularForms(),
		OfflineForms:     r.GetOfflineForms(),
		AnswerKeys:       r.GetAnswerKeys(),
		SyllabusList:     r.GetSyllabus(),
		SarkariYojnaList: r.GetSarkariYojna(),
		VerificationList: r.GetVerification(),
		UpcomingList:     r.GetUpcoming(),
	}
}

func (r *Repository) InsertJob(job models.Job) error {
	_, err := r.DB.Exec(`
		INSERT INTO jobs (title, description, url, last_date) 
		VALUES ($1, $2, $3, $4)`,
		job.Title, job.Description, job.Url, job.LastDate,
	)
	return err
}

func (r *Repository) InsertResult(res models.Result) error {
	_, err := r.DB.Exec(`
		INSERT INTO results (title, url, exam_date) 
		VALUES ($1, $2, $3)`,
		res.Title, res.Url, res.ExamDate,
	)
	return err
}

func (r *Repository) InsertAdmitCard(admit models.AdmitCard) error {
	_, err := r.DB.Exec(`
		INSERT INTO admit_cards (title, url, exam_date) 
		VALUES ($1, $2, $3)`,
		admit.Title, admit.Url, admit.ExamDate,
	)
	return err
}

func (r *Repository) InsertVimpLink(link models.VImpLink) error {
	_, err := r.DB.Exec(`
		INSERT INTO vimps_links (title, url) 
		VALUES ($1, $2)`,
		link.Title, link.Url,
	)
	return err
}

func (r *Repository) InsertImpLink(link models.ImpLink) error {
	_, err := r.DB.Exec(`
		INSERT INTO imp_links (title, url) 
		VALUES ($1, $2)`,
		link.Title, link.Url,
	)
	return err
}

func (r *Repository) InsertNaukriForm(f models.NaukriForm) error {
	_, err := r.DB.Exec(`INSERT INTO naukri_forms (title, url) VALUES ($1, $2)`, f.Title, f.Url)
	return err
}

func (r *Repository) InsertAdmission(a models.Admission) error {
	_, err := r.DB.Exec(`INSERT INTO admissions (title, url) VALUES ($1, $2)`, a.Title, a.Url)
	return err
}

func (r *Repository) InsertRegularForm(f models.RegularForm) error {
	_, err := r.DB.Exec(`INSERT INTO regular_forms (title, url) VALUES ($1, $2)`, f.Title, f.Url)
	return err
}

func (r *Repository) InsertOfflineForm(f models.OfflineForm) error {
	_, err := r.DB.Exec(`INSERT INTO offline_forms (title, url) VALUES ($1, $2)`, f.Title, f.Url)
	return err
}

func (r *Repository) InsertAnswerKey(a models.AnswerKey) error {
	_, err := r.DB.Exec(`INSERT INTO answer_keys (title, url) VALUES ($1, $2)`, a.Title, a.Url)
	return err
}

func (r *Repository) InsertSyllabus(s models.Syllabus) error {
	_, err := r.DB.Exec(`INSERT INTO syllabus (title, url) VALUES ($1, $2)`, s.Title, s.Url)
	return err
}

func (r *Repository) InsertSarkariYojna(s models.SarkariYojna) error {
	_, err := r.DB.Exec(`INSERT INTO sarkari_yojna (title, url) VALUES ($1, $2)`, s.Title, s.Url)
	return err
}

func (r *Repository) InsertVerification(v models.Verification) error {
	_, err := r.DB.Exec(`INSERT INTO verification (title, url) VALUES ($1, $2)`, v.Title, v.Url)
	return err
}

func (r *Repository) InsertUpcoming(u models.Upcoming) error {
	_, err := r.DB.Exec(`INSERT INTO upcoming (title, url) VALUES ($1, $2)`, u.Title, u.Url)
	return err
}

func (r *Repository) InsertAllData(data models.AllData) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Step 1: Clear all tables (TRUNCATE resets serial IDs)
	_, err = tx.Exec(`
		TRUNCATE jobs, results, admit_cards, vimps_links, imp_links,
		         naukri_forms, admissions, regular_forms, offline_forms,
		         answer_keys, syllabus, sarkari_yojna, verification, upcoming
		RESTART IDENTITY CASCADE`)
	if err != nil {
		return err
	}

	// Insert jobs
	for _, job := range data.Jobs {
		_, err := tx.Exec(`
			INSERT INTO jobs (title, description, url, last_date) 
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
			INSERT INTO results (title, url, exam_date) 
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
			INSERT INTO admit_cards (title, url, exam_date) 
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
			INSERT INTO vimps_links (title, url) 
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
			INSERT INTO imp_links (title, url) 
			VALUES ($1, $2)`,
			ilink.Title, ilink.Url,
		)
		if err != nil {
			return err
		}
	}

	// Insert naukri_forms
	for _, form := range data.NaukriForms {
		_, err := tx.Exec(`INSERT INTO naukri_forms (title, url) VALUES ($1, $2)`,
			form.Title, form.Url)
		if err != nil {
			return err
		}
	}

	// Insert admissions
	for _, adm := range data.Admissions {
		_, err := tx.Exec(`INSERT INTO admissions (title, url) VALUES ($1, $2)`,
			adm.Title, adm.Url)
		if err != nil {
			return err
		}
	}

	// Insert regular_forms
	for _, form := range data.RegularForms {
		_, err := tx.Exec(`INSERT INTO regular_forms (title, url) VALUES ($1, $2)`,
			form.Title, form.Url)
		if err != nil {
			return err
		}
	}

	// Insert offline_forms
	for _, form := range data.OfflineForms {
		_, err := tx.Exec(`INSERT INTO offline_forms (title, url) VALUES ($1, $2)`,
			form.Title, form.Url)
		if err != nil {
			return err
		}
	}

	// Insert answer_keys
	for _, ans := range data.AnswerKeys {
		_, err := tx.Exec(`INSERT INTO answer_keys (title, url) VALUES ($1, $2)`,
			ans.Title, ans.Url)
		if err != nil {
			return err
		}
	}

	// Insert syllabus
	for _, s := range data.SyllabusList {
		_, err := tx.Exec(`INSERT INTO syllabus (title, url) VALUES ($1, $2)`,
			s.Title, s.Url)
		if err != nil {
			return err
		}
	}

	// Insert sarkari_yojna
	for _, sy := range data.SarkariYojnaList {
		_, err := tx.Exec(`INSERT INTO sarkari_yojna (title, url) VALUES ($1, $2)`,
			sy.Title, sy.Url)
		if err != nil {
			return err
		}
	}

	// Insert verification
	for _, v := range data.VerificationList {
		_, err := tx.Exec(`INSERT INTO verification (title, url) VALUES ($1, $2)`,
			v.Title, v.Url)
		if err != nil {
			return err
		}
	}

	// Insert upcoming
	for _, u := range data.UpcomingList {
		_, err := tx.Exec(`INSERT INTO upcoming (title, url) VALUES ($1, $2)`,
			u.Title, u.Url)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

