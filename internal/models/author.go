package models

import "github.com/alexalreadytaken/book-service/proto"

type Author struct {
	Id      int64  `db:"id"`
	Name    string `db:"name"`
	Surname string `db:"surname"`
}

func (a *Author) ToProto() *proto.Author {
	return &proto.Author{
		Id:      a.Id,
		Name:    a.Name,
		Surname: a.Surname,
	}
}
