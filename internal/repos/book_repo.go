package repos

import (
	"fmt"

	"github.com/alexalreadytaken/book-service/internal/models"
	"github.com/alexalreadytaken/book-service/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//usually use statements
//db for extreme sutations
type MysqlBookRepo struct {
	db              *sqlx.DB
	bookAuthorsStmt *sqlx.Stmt
	authorBooksStmt *sqlx.Stmt
}

//mysql queries for repo statements
const (
	bookAuthorsQuery string = `
		select a.id, a.name, a.surname from book_author as ba 
		inner join author as a on ba.author_id=a.id where ba.book_id=?`
	authorBooksQuery string = `
		select b.id, b.name, b.pages_count from book_author as ba 
		inner join book as b on ba.book_id=b.id where ba.author_id=?`
)

//try connect to db
//validate & preapare statements
func NewMysqlBookRepo(cnf *utils.AppConfig) (*MysqlBookRepo, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cnf.DbUser, cnf.DbPwd, cnf.DbAddr, cnf.DbPort, cnf.DbSchema)
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		return nil, err
	}
	bookAuthorsStmt, err := db.Preparex(bookAuthorsQuery)
	if err != nil {
		return nil, err
	}
	authorBooksStmt, err := db.Preparex(authorBooksQuery)
	if err != nil {
		return nil, err
	}
	return &MysqlBookRepo{
		db:              db,
		bookAuthorsStmt: bookAuthorsStmt,
		authorBooksStmt: authorBooksStmt,
	}, nil
}

//execute statement and map result to array
func (repo *MysqlBookRepo) BookAuthors(bookId int64) ([]models.Author, error) {
	var authors []models.Author
	err := repo.bookAuthorsStmt.Select(&authors, bookId)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

//execute statement and map result to array
func (repo *MysqlBookRepo) AuthorBooks(authorId int64) ([]models.Book, error) {
	var books []models.Book
	err := repo.authorBooksStmt.Select(&books, authorId)
	if err != nil {
		return nil, err
	}
	return books, nil
}
