package mariadb

import (
	"database/sql"
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

func (repo *mariaRepository) InsertBook(book *models.Book) error {
	_, err := repo.db.Exec(`INSERT INTO books (id, bookName, authorName, publishDate) VALUES(UNHEX(?), ?, ?, DATE(?))`, uuid.New(), book.BookName, book.AuthorName, book.PublishDate.String())
	return err
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
	var book *models.Book
	rows, err := repo.db.Query(`SELECT lower(HEX(id)), bookName, authorName, publishDate FROM books where id=UNHEX(?)`, bookId)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			book = &models.Book{}
			err = rows.Scan(&book.ID, &book.BookName, &book.AuthorName, &book.PublishDate)
			if err != nil {
				return nil, err
			}
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return book, nil
}
func (repo *mariaRepository) UpdateBook(bookId string, book *models.Book) error {
	_, err := repo.db.Exec(`UPDATE books SET bookName=?, authorName=?, publishDate=DATE(?) WHERE id=UNHEX(?)`, book.BookName, book.AuthorName, book.PublishDate.String(), bookId)
	return err
}
func (repo *mariaRepository) DeleteBook(bookId string) error {
	_, err := repo.db.Exec(`DELETE FROM books WHERE id=UNHEX(?)`, bookId)
	return err
}


func (repo *mariaRepository) InsertReviewer(reviewer *models.Reviewer) error {
	_, err := repo.db.Exec(`INSERT INTO reviewers (id, name) VALUES(UNHEX(?), ?)`, uuid.New(), reviewer.Name)
	return err
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
