package models

import "github.com/alexalreadytaken/book-service/proto"

type Book struct {
	Id         int64  `db:"id"`
	Name       string `db:"name"`
	PagesCount int64  `db:"pages_count"`
}

func (b *Book) ToProto() *proto.Book {
	return &proto.Book{
		Id:         b.Id,
		Name:       b.Name,
		PagesCount: b.PagesCount,
	}
}
