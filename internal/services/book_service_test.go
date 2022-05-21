package services

import (
	"context"
	"net"
	"testing"

	"github.com/alexalreadytaken/book-service/internal/repos"
	"github.com/alexalreadytaken/book-service/internal/utils"
	"github.com/alexalreadytaken/book-service/proto"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type BookServiceTestSuite struct {
	suite.Suite
	bookServiceClient proto.BookServiceClient
}

func TestBookService(t *testing.T) {
	suite.Run(t, &BookServiceTestSuite{})
}

func (s *BookServiceTestSuite) SetupSuite() {
	lis := s.upServer()
	s.upClient(lis)
}

func (s *BookServiceTestSuite) TestAuthorBooksFromInitSQL() {
	req := &proto.AuthorBooksRequest{AuthorId: 1}
	resp, err := s.bookServiceClient.AuthorBooks(context.Background(), req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(2, len(resp.Books))
	for _, book := range resp.Books {
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
	req = &proto.AuthorBooksRequest{AuthorId: 2}
	resp, err = s.bookServiceClient.AuthorBooks(context.Background(), req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, len(resp.Books))
	book := resp.Books[0]
	s.Equal(int64(2), book.Id)
	s.Equal("Beyond Games and Scripts", book.Name)
	s.Equal(int64(500), book.PagesCount)
}

func (s *BookServiceTestSuite) TestBookAuthorsFromInitSQL() {
	req := &proto.BookAuthorsRequest{BookId: 1}
	resp, err := s.bookServiceClient.BookAuthors(context.Background(), req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, len(resp.Authors))
	author := resp.Authors[0]
	s.Equal(int64(1), author.Id)
	s.Equal("Eric", author.Name)
	s.Equal("Berne", author.Surname)
	req = &proto.BookAuthorsRequest{BookId: 2}
	resp, err = s.bookServiceClient.BookAuthors(context.Background(), req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(2, len(resp.Authors))
	for _, author := range resp.Authors {
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

func (s *BookServiceTestSuite) upServer() *bufconn.Listener {
	cnf := utils.LoadCnfFromEnv()
	lis := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	repo, err := repos.NewMysqlBookRepo(cnf)
	if err != nil {
		s.Fail("can`t init book repo from env config=", err)
	}
	service := NewBookService(repo)
	proto.RegisterBookServiceServer(server, service)
	go func() {
		if err := server.Serve(lis); err != nil {
			s.Fail(err.Error())
		}
	}()
	return lis
}

func (s *BookServiceTestSuite) upClient(lis *bufconn.Listener) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				return lis.Dial()
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.Fail("can`t up grpc dial context=", err.Error())
	}
	client := proto.NewBookServiceClient(conn)
	s.bookServiceClient = client
}
