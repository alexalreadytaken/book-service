package services

import (
	"context"
	"log"

	"github.com/alexalreadytaken/book-service/internal/repos"
	"github.com/alexalreadytaken/book-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookService struct {
	bookRepo *repos.MysqlBookRepo
}

func NewBookService(bookRepo *repos.MysqlBookRepo) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

//get book authors from repo and map them to proto structure
func (service *BookService) BookAuthors(
	ctx context.Context,
	req *proto.BookAuthorsRequest) (*proto.BookAuthorsResponse, error) {
	authors, err := service.bookRepo.BookAuthors(req.BookId)
	if err != nil {
		log.Println("error while gettinf info about book authors=", err.Error())
		return nil, status.Errorf(codes.Internal, "can`t get book authors")
	}
	protoAuthors := make([]*proto.Author, len(authors))
	for i := 0; i < len(authors); i++ {
		protoAuthors[i] = authors[i].ToProto()
	}
	return &proto.BookAuthorsResponse{
		Authors: protoAuthors,
	}, nil
}

//get author books from repo and map them to proto structure
func (service *BookService) AuthorBooks(
	ctx context.Context,
	req *proto.AuthorBooksRequest) (*proto.AuthorBooksResponse, error) {
	books, err := service.bookRepo.AuthorBooks(req.AuthorId)
	if err != nil {
		log.Println("error while gettinf info about author books=", err.Error())
		return nil, status.Errorf(codes.Internal, "can`t get author books")
	}
	protoBooks := make([]*proto.Book, len(books))
	for i := 0; i < len(books); i++ {
		protoBooks[i] = books[i].ToProto()
	}
	return &proto.AuthorBooksResponse{
		Books: protoBooks,
	}, nil
}
