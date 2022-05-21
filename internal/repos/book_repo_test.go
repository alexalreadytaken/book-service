package repos

import (
	"testing"

	"github.com/alexalreadytaken/book-service/internal/utils"
	"github.com/stretchr/testify/suite"
)

type BookRepoTestSuite struct {
	suite.Suite
	repo *MysqlBookRepo
}

func TestBookRepo(t *testing.T) {
	suite.Run(t, &BookRepoTestSuite{})
}

//load config from env and init repo
func (s *BookRepoTestSuite) SetupSuite() {
	cnf := utils.LoadCnfFromEnv()
	repo, err := NewMysqlBookRepo(cnf)
	if err != nil {
		s.Fail("can`t init repo with env", err)
	}
	s.repo = repo
}

//get the authors books and check them according to the init.sql file
func (s *BookRepoTestSuite) TestAuthorBooksFromInitSQL() {
	s.T().Log("Start testing relation author to book from init.sql")
	books, err := s.repo.AuthorBooks(1)
	s.NoError(err)
	s.NotNil(books)
	s.Equal(2, len(books))
	for _, book := range books {
		switch book.Id {
		case 1:
			s.Equal("Games People Play", book.Name)
			s.Equal(int64(563), book.PagesCount)
		case 2:
			s.Equal("Beyond Games and Scripts", book.Name)
			s.Equal(int64(500), book.PagesCount)
		default:
			s.Fail("unexpected book=", book)
		}
	}
	books, err = s.repo.AuthorBooks(2)
	s.NoError(err)
	s.NotNil(books)
	s.Equal(1, len(books))
	book := books[0]
	s.Equal(int64(2), book.Id)
	s.Equal("Beyond Games and Scripts", book.Name)
	s.Equal(int64(500), book.PagesCount)
}

//get the books authors and check them according to the init.sql file
func (s *BookRepoTestSuite) TestBookAuthorsFromInitSQL() {
	authors, err := s.repo.BookAuthors(1)
	s.NoError(err)
	s.NotNil(authors)
	s.Equal(1, len(authors))
	author := authors[0]
	s.Equal(int64(1), author.Id)
	s.Equal("Eric", author.Name)
	s.Equal("Berne", author.Surname)
	authors, err = s.repo.BookAuthors(2)
	s.NoError(err)
	s.NotNil(authors)
	s.Equal(2, len(authors))
	for _, author := range authors {
		switch author.Id {
		case 1:
			s.Equal("Eric", author.Name)
			s.Equal("Berne", author.Surname)
		case 2:
			s.Equal("*more*", author.Name)
			s.Equal("*author*", author.Surname)
		default:
			s.Fail("unexpected author=", author)
		}
	}
}
