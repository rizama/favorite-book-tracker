package entity

import "time"

type Book struct {
	Title     string
	Author    string
	Rating    int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
