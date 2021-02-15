package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/Sumeet-Ranjan-Parida/ContactBook/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Failed to listen on port 4040: %v", err)
	}

	gs := grpc.NewServer()

	if err := gs.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server on port 4040: %v", err)
	}
}

func (s *server) Getcontact(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	db, err := sql.Open("mysql", "root:sumeet@tcp(127.0.0.1:3306)/contactbook")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	name, number := request.GetName(), request.GetNumber()

	insert, err := db.Prepare("INSERT INTO contacts(name, phno) VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, number)

	return &proto.Response{Cname: name, Cnumber: number}, nil
}
