package models

type Repository interface {
	InsertBook(book *Book) error
	GetAllBooks() ([]*Book, error)
	GetBook(bookId string) (*Book, error)
	UpdateBook(bookId string, book *Book) error
	DeleteBook(bookId string) error

	InsertReviewer(reviewer *Reviewer) error
	GetAllReviewers() ([]*Reviewer, error)
	GetReviewer(reviewerId string) (*Reviewer, error)
	UpdateReviewer(reviewerId string, reviewer *Reviewer) error
	DeleteReviewer(reviewerId string) error
}
