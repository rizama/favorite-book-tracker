package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	// Create a sample Book instance
	book := Book{
		Title:     "Sample Book",
		Author:    "John Doe",
		Rating:    4,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Perform assertions
	assert.Equal(t, "Sample Book", book.Title, "Title should be 'Sample Book'")
	assert.Equal(t, "John Doe", book.Author, "Author should be 'John Doe'")
	assert.Equal(t, int8(4), book.Rating, "Rating should be 4")

	// You can perform more assertions based on your needs
}

// Add more test functions as needed
