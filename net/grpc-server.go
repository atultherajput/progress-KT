package net

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/atultherajput/go_crash_course/database/grom/dao"
	"github.com/atultherajput/go_crash_course/database/grom/dbinit"
	"github.com/atultherajput/go_crash_course/grpc/protos"
	"github.com/atultherajput/go_crash_course/models"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Handler dao.Handler

type server struct {
	protos.UnimplementedBooksServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetAllBooks(ctx context.Context, in *protos.NoParamRequest) (*protos.BooksResponse, error) {
	log.Printf("Received: Get all books request")

	//Find all books
	books := Handler.GetAll()

	var bookArray []*protos.BookRequestResponse

	serailizedBytes, _ := json.Marshal(books)
	if err := json.Unmarshal(serailizedBytes, &bookArray); err != nil {
		log.Printf("Error in parsing getAll")
	}

	return &protos.BooksResponse{Books: bookArray}, nil
}

func (s *server) GetBook(ctx context.Context, in *protos.BookIdRequest) (*protos.BookRequestResponse, error) {
	log.Printf("Received: Get a book request")

	id := in.GetId()

	// Find book by Id
	book := Handler.Get(id)

	return &protos.BookRequestResponse{Title: book.Title, Author: book.Author, Desc: book.Desc}, nil
}

func (s *server) AddBook(ctx context.Context, in *protos.BookRequestResponse) (*protos.TextResponse, error) {
	log.Printf("Received: Add a book request")

	var book models.Book
	book.Title = in.GetTitle()
	book.Author = in.GetAuthor()
	book.Desc = in.GetDesc()

	// Append to the Books table
	Handler.Add(book)

	return &protos.TextResponse{Message: "The book has been inserted successfully!"}, nil
}

func (s *server) DeleteBook(ctx context.Context, in *protos.BookIdRequest) (*protos.TextResponse, error) {
	log.Printf("Received: Delete a book request")

	id := in.GetId()

	// Delete that book
	Handler.Delete(id)

	return &protos.TextResponse{Message: "The book has been deleted successfully!"}, nil
}

func (s *server) UpdateBook(ctx context.Context, in *protos.BookUpdateRequest) (*protos.TextResponse, error) {
	log.Printf("Received: Update a book request")

	id := in.GetId()

	var updatedBook models.Book
	updatedBook.Title = in.Book.GetTitle()
	updatedBook.Author = in.Book.GetAuthor()
	updatedBook.Desc = in.Book.GetDesc()

	//Update DB
	Handler.Update(id, updatedBook)

	return &protos.TextResponse{Message: "The book has been upadted successfully!"}, nil
}

const (
	grpcServerEndpoint = "localhost:%s"
)

func RunGrpc(port *int) {

	DB := dbinit.Init()
	Handler = dao.New(DB)

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Books service to the server
	protos.RegisterBooksServer(s, &server{})

	// Serve gRPC server
	log.Printf("serving gRPC on %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	//Register REST handler
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := protos.RegisterBooksHandlerFromEndpoint(ctx, mux, fmt.Sprintf(grpcServerEndpoint, "50051"), opts); err != nil {
		log.Fatalf("Failed to register gRPC gateway service endpoint: %v", err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux); err != nil {
		log.Fatalf("Could not setup HTTP endpoint: %v", err)
	}

}
