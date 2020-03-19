package db

type Script struct {
	username  string
	auth_id   string
	script_id string
	ratings   []Ratings
	comments  []Comments
}
