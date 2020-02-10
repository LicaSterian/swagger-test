package mariadb

import (
	"database/sql"
	"fmt"
	"github.com/LicaSterian/swagger-test/models"
	"github.com/brideclick/uuid"
	"github.com/spf13/viper"

	"github.com/brideclick/initialize"

	_ "github.com/go-sql-driver/mysql"
)

type mariaRepository struct {
	db *sql.DB
}

func init() {
	initialize.Viper()
}

func NewMariaRepository() (models.Repository, error) {
	db, err := sql.Open("mysql", viper.GetString("mariaDB"))
	if err != nil {
		return nil, err
	}

	repo := &mariaRepository{
		db: db,
	}
	return repo, nil
}

func (repo *mariaRepository) InsertBook(book *models.Book) (*models.Book, error) {
	book.ID = uuid.New()
	_, err := repo.db.Exec(`INSERT INTO books (id, bookName, authorName, publishDate) VALUES(UNHEX(?), ?, ?, DATE(?))`, book.ID, book.BookName, book.AuthorName, book.PublishDate.String())
	return book, err
}
func (repo *mariaRepository) GetAllBooks() ([]*models.Book, error) {
	var result []*models.Book
	rows, err := repo.db.Query(`SELECT lower(HEX(id)), bookName, authorName, publishDate FROM books`)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		for rows.Next() {
			book := &models.Book{}
			err = rows.Scan(&book.ID, &book.BookName, &book.AuthorName, &book.PublishDate)
			if err != nil {
				return nil, err
			}
			result = append(result, book)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (repo *mariaRepository) GetBook(bookId string) (*models.Book, error) {
	book := &models.Book{}
	err := repo.db.QueryRow(`SELECT lower(HEX(id)), bookName, authorName, publishDate FROM books where id=UNHEX(?)`, bookId).
		Scan(&book.ID, &book.BookName, &book.AuthorName, &book.PublishDate)
	if err != nil {
		return nil, err
	}
	reviewers, err := repo.GetReviewersByBookId(bookId)
	if err != nil {
		return nil, err
	}
	book.Reviewers = reviewers
	return book, nil
}
func (repo *mariaRepository) UpdateBook(bookId string, book *models.Book) error {
	oldReviewers, err := repo.GetReviewersByBookId(bookId)
	if err != nil {
		return err
	}

	// insert
	for _, n := range book.Reviewers {
		if !containsReviewer(oldReviewers, n) {
			err = repo.insertBookReviewer(bookId, n.ID)
			if err != nil {
				return fmt.Errorf("repo.insertBookReviewer error: %s", err.Error())
			}
		}
	}
	// delete
	for _, o := range oldReviewers {
		if !containsReviewer(book.Reviewers, o) {
			err = repo.deleteBookReviewer(bookId, o.ID)
			if err != nil {
				return fmt.Errorf("repo.deleteBookReviewer error: %s", err.Error())
			}
		}
	}

	_, err = repo.db.Exec(`UPDATE books SET bookName=?, authorName=?, publishDate=DATE(?) WHERE id=UNHEX(?)`,
		book.BookName, book.AuthorName, book.PublishDate.String(), bookId)
	return err
}
func (repo *mariaRepository) DeleteBook(bookId string) error {
	_, err := repo.db.Exec(`DELETE FROM books WHERE id=UNHEX(?)`, bookId)
	return err
}


func (repo *mariaRepository) InsertReviewer(reviewer *models.Reviewer) (*models.Reviewer, error) {
	reviewer.ID = uuid.New()
	_, err := repo.db.Exec(`INSERT INTO reviewers (id, name) VALUES(UNHEX(?), ?)`, reviewer.ID, reviewer.Name)
	return reviewer, err
}
func (repo *mariaRepository) GetAllReviewers() ([]*models.Reviewer, error) {
	var result []*models.Reviewer
	rows, err := repo.db.Query(`SELECT lower(HEX(id)), name FROM reviewers`)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		for rows.Next() {
			reviewer := &models.Reviewer{}
			err = rows.Scan(&reviewer.ID, &reviewer.Name)
			if err != nil {
				return nil, err
			}
			result = append(result, reviewer)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (repo *mariaRepository) GetReviewer(reviewerId string) (*models.Reviewer, error) {
	reviewer := &models.Reviewer{}
	err := repo.db.QueryRow(`SELECT lower(HEX(id)), name FROM reviewers where id=UNHEX(?)`, reviewerId).Scan(&reviewer.ID, &reviewer.Name)
	if err != nil {
		return nil, err
	}
	return reviewer, nil
}
func (repo *mariaRepository) UpdateReviewer(reviewerId string, reviewer *models.Reviewer) error {
	_, err := repo.db.Exec(`UPDATE reviewers SET name=? WHERE id=UNHEX(?)`, reviewer.Name, reviewerId)
	return err
}
func (repo *mariaRepository) DeleteReviewer(reviewerId string) error {
	_, err := repo.db.Exec(`DELETE FROM reviewers WHERE id=UNHEX(?)`, reviewerId)
	return err
}



func (repo *mariaRepository) GetReviewersByBookId(bookId string) ([]*models.Reviewer, error) {
	var result []*models.Reviewer
	rows, err := repo.db.Query(`SELECT lower(HEX(id)), name FROM reviewers r LEFT JOIN books_reviewers br ON r.id = br.reviewerId WHERE br.bookId=UNHEX(?)`, bookId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		for rows.Next() {
			reviewer := &models.Reviewer{}
			err = rows.Scan(&reviewer.ID, &reviewer.Name)
			if err != nil {
				return nil, err
			}
			result = append(result, reviewer)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (repo *mariaRepository) insertBookReviewer(bookId, reviewerId string) error {
	_, err := repo.db.Exec(`INSERT INTO books_reviewers (bookId, reviewerId) VALUES(UNHEX(?), UNHEX(?))`,
		bookId, reviewerId)
	return err
}

func (repo *mariaRepository) deleteBookReviewer(bookId, reviewerId string) error {
	_, err := repo.db.Exec(`DELETE FROM books_reviewers WHERE bookId=UNHEX(?) AND reviewerId=UNHEX(?)`,
		bookId, reviewerId)
	return err
}

func containsBook(s []*models.Book, e *models.Book) bool {
	for _, v := range s {
		if v.ID == e.ID {
			return true
		}
	}
	return false
}
func containsReviewer(s []*models.Reviewer, e *models.Reviewer) bool {
	for _, v := range s {
		if v.ID == e.ID {
			return true
		}
	}
	return false
}