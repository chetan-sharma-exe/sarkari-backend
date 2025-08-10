package models

import "time"

type Job struct {
	ID          int        `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	ApplyURL    string     `db:"apply_url" json:"apply_url"`
	LastDate    *time.Time `db:"last_date" json:"last_date,omitempty"`
	DatePosted  time.Time  `db:"date_posted" json:"date_posted"`
}

type Result struct {
	ID         int        `db:"id" json:"id"`
	Title      string     `db:"title" json:"title"`
	ResultLink string     `db:"result_link" json:"result_link"`
	ExamDate   *time.Time `db:"exam_date" json:"exam_date,omitempty"`
	DatePosted time.Time  `db:"date_posted" json:"date_posted"`
}

type AdmitCard struct {
	ID           int        `db:"id" json:"id"`
	Title        string     `db:"title" json:"title"`
	DownloadLink string     `db:"download_link" json:"download_link"`
	ExamDate     *time.Time `db:"exam_date" json:"exam_date,omitempty"`
	DatePosted   time.Time  `db:"date_posted" json:"date_posted"`
}

type VImpLink struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Link  string `db:"link" json:"link"`
}

type ImpLink struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Link  string `db:"link" json:"link"`
}
