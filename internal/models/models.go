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


type NaukriForm struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type Admission struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type RegularForm struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type OfflineForm struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type AnswerKey struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type Syllabus struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type SarkariYojna struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type Verification struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}


type Upcoming struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Url   string `db:"url" json:"url"`
}

type AllData struct {
	Jobs             []Job          `json:"jobs"`
	Results          []Result       `json:"results"`
	AdmitCards       []AdmitCard    `json:"admit_cards"`
	VimpLinks        []VImpLink     `json:"vimps_links"`
	ImpLinks         []ImpLink      `json:"imp_links"`
	NaukriForms      []NaukriForm   `json:"naukri_forms"`
	Admissions       []Admission    `json:"admissions"`
	RegularForms     []RegularForm  `json:"regular_forms"`
	OfflineForms     []OfflineForm  `json:"offline_forms"`
	AnswerKeys       []AnswerKey    `json:"answer_keys"`
	SyllabusList     []Syllabus     `json:"syllabus"`
	SarkariYojnaList []SarkariYojna `json:"sarkari_yojna"`
	VerificationList []Verification `json:"verification"`
	UpcomingList     []Upcoming     `json:"upcoming"`
}
