package svc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Neakxs/protocel-example/internal/db"
	v1 "github.com/Neakxs/protocel-example/proto/library/v1"
	"github.com/google/uuid"
)

func NewLibraryService(wu db.WorkUnit) LibraryService {
	return &libraryService{wu: wu}
}

type LibraryService interface {
	CreateAuthor(ctx context.Context, in *v1.CreateAuthorRequest) (*v1.Author, error)
	ListAuthors(ctx context.Context, in *v1.ListAuthorsRequest) (*v1.ListAuthorsResponse, error)
	DeleteAuthor(ctx context.Context, in *v1.DeleteAuthorRequest) (*v1.Author, error)
	CreateBook(ctx context.Context, in *v1.CreateBookRequest) (*v1.Book, error)
	ListBooks(ctx context.Context, in *v1.ListBooksRequest) (*v1.ListBooksResponse, error)
	DeleteBook(ctx context.Context, in *v1.DeleteBookRequest) (*v1.Book, error)
}

type libraryService struct {
	wu db.WorkUnit
}

func (s *libraryService) CreateAuthor(ctx context.Context, in *v1.CreateAuthorRequest) (*v1.Author, error) {
	if err := in.Validate(ctx); err != nil {
		return nil, err
	}
	in.Author.Name = fmt.Sprintf("authors/%s", uuid.NewString())
	if err := s.wu.Authors().SaveAuthor(ctx, in.Author); err != nil {
		return nil, err
	}
	return in.Author, nil
}

func (s *libraryService) ListAuthors(ctx context.Context, in *v1.ListAuthorsRequest) (*v1.ListAuthorsResponse, error) {
	if err := in.Validate(ctx); err != nil {
		return nil, err
	}
	if in.PageSize == 0 {
		in.PageSize = 20
	}
	pageOffset := 0
	if in.PageToken != "" {
		if i, err := strconv.ParseInt(in.PageToken, 16, 64); err != nil {
			return nil, err
		} else {
			pageOffset = int(i)
		}
	}
	if as, err := s.wu.Authors().QueryAuthors(ctx, int(in.PageSize), pageOffset); err != nil {
		return nil, err
	} else {
		return &v1.ListAuthorsResponse{
			Authors:       as,
			NextPageToken: strconv.FormatInt(int64(pageOffset)+int64(in.PageSize), 16),
		}, nil
	}
}

func (s *libraryService) DeleteAuthor(ctx context.Context, in *v1.DeleteAuthorRequest) (*v1.Author, error) {
	if err := in.Validate(ctx); err != nil {
		return nil, err
	}
	if a, err := s.wu.Authors().DeleteAuthor(ctx, in.Name); err != nil {
		return nil, err
	} else {
		return a, nil
	}
}

func (s *libraryService) CreateBook(ctx context.Context, in *v1.CreateBookRequest) (*v1.Book, error) {
	if err := in.Validate(ctx); err != nil {
		return nil, err
	}
	in.Book.Name = fmt.Sprintf("%s/books/%s", in.Parent, uuid.NewString())
	if err := s.wu.Books().SaveBook(ctx, in.Book); err != nil {
		return nil, err
	}
	return in.Book, nil
}

func (s *libraryService) ListBooks(ctx context.Context, in *v1.ListBooksRequest) (*v1.ListBooksResponse, error) {
	if err := in.Validate(ctx); err != nil {
		return nil, err
	}
	if in.PageSize == 0 {
		in.PageSize = 20
	}
	pageOffset := 0
	if in.PageToken != "" {
		if i, err := strconv.ParseInt(in.PageToken, 16, 64); err != nil {
			return nil, err
		} else {
			pageOffset = int(i)
		}
	}
	if bs, err := s.wu.Books().QueryBooks(ctx, int(in.PageSize), pageOffset); err != nil {
		return nil, err
	} else {
		return &v1.ListBooksResponse{
			Books:         bs,
			NextPageToken: strconv.FormatInt(int64(pageOffset)+int64(in.PageSize), 16),
		}, nil
	}
}

func (s *libraryService) DeleteBook(ctx context.Context, in *v1.DeleteBookRequest) (*v1.Book, error) {
	if err := in.Validate(ctx); err != nil {
		return nil, err
	}
	if b, err := s.wu.Books().DeleteBook(ctx, in.Name); err != nil {
		return nil, err
	} else {
		return b, nil
	}
}
