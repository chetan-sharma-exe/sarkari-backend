package models

type Job struct {
	ID          int     `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Url         string  `db:"url" json:"url"`
	LastDate    *string `db:"last_date" json:"last_date,omitempty"`
	DatePosted  string  `db:"date_posted" json:"date_posted"`
}

type Result struct {
	ID         int     `db:"id" json:"id"`
	Title      string  `db:"title" json:"title"`
	Url        string  `db:"url" json:"url"`
	ExamDate   *string `db:"exam_date" json:"exam_date,omitempty"`
	DatePosted string  `db:"date_posted" json:"date_posted"`
}

type AdmitCard struct {
	ID         int     `db:"id" json:"id"`
	Title      string  `db:"title" json:"title"`
	Url        string  `db:"url" json:"url"`
	ExamDate   *string `db:"exam_date" json:"exam_date,omitempty"`
	DatePosted string  `db:"date_posted" json:"date_posted"`
}

type VImpLink struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}

type ImpLink struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}

type AllData struct {
	Jobs       []Job       `json:"jobs"`
	Results    []Result    `json:"results"`
	AdmitCards []AdmitCard `json:"admit_cards"`
	VimpLinks  []VImpLink  `json:"vimps_links"`
	ImpLinks   []ImpLink   `json:"imp_links"`
}
