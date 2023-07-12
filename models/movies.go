package models

// data models goes here
type Movie struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	ReleaseYear int      `json:"releaseYear"`
	Artists     []string `json:"artists"`
}
