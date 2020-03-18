package db

type Metadata struct {
	username string
	auth_id   string
	scripts []Scripts
	ratings []Ratings
	comments []Comments
}