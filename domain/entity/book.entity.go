package entity

import "time"

type Book struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Rating    int8       `json:"rating"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
