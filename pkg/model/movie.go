package model

import "time"

// Movie model
type Movie struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    ReleaseDate time.Time `json:"release_date"`
}
