package main

import (
	"fmt"
	"log"
	"net"

	"github.com/alexalreadytaken/book-service/internal/repos"
	"github.com/alexalreadytaken/book-service/internal/services"
	"github.com/alexalreadytaken/book-service/internal/utils"
	"github.com/alexalreadytaken/book-service/proto"
	"google.golang.org/grpc"
)

//load config from env
//init serivces and register them to grpc server
//start server
func main() {
	cnf := utils.LoadCnfFromEnv() 
	bookRepo, err := repos.NewMysqlBookRepo(cnf)
	if err != nil {
		log.Fatal(err.Error())
	}
	bookService := services.NewBookService(bookRepo)
	server := grpc.NewServer()
	proto.RegisterBookServiceServer(server, bookService)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cnf.SerivcePort))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("starting grpc server")
	log.Fatal(server.Serve(listener))
}
