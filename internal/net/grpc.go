package net

import (
	"context"

	"github.com/Neakxs/protocel-example/internal/svc"
	v1 "github.com/Neakxs/protocel-example/proto/library/v1"
)

func NewLibraryService(s svc.LibraryService) v1.LibraryServiceServer {
	return &libraryServiceServer{s: s}
}

type libraryServiceServer struct {
	s svc.LibraryService
	*v1.UnimplementedLibraryServiceServer
}

func (s *libraryServiceServer) CreateAuthor(ctx context.Context, in *v1.CreateAuthorRequest) (*v1.Author, error) {
	return s.s.CreateAuthor(ctx, in)
}
func (s *libraryServiceServer) ListAuthors(ctx context.Context, in *v1.ListAuthorsRequest) (*v1.ListAuthorsResponse, error) {
	return s.s.ListAuthors(ctx, in)
}
func (s *libraryServiceServer) DeleteAuthor(ctx context.Context, in *v1.DeleteAuthorRequest) (*v1.Author, error) {
	return s.s.DeleteAuthor(ctx, in)
}
func (s *libraryServiceServer) CreateBook(ctx context.Context, in *v1.CreateBookRequest) (*v1.Book, error) {
	return s.s.CreateBook(ctx, in)
}
func (s *libraryServiceServer) ListBooks(ctx context.Context, in *v1.ListBooksRequest) (*v1.ListBooksResponse, error) {
	return s.s.ListBooks(ctx, in)
}
func (s *libraryServiceServer) DeleteBook(ctx context.Context, in *v1.DeleteBookRequest) (*v1.Book, error) {
	return s.s.DeleteBook(ctx, in)
}
