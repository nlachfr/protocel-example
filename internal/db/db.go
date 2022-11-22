package db

import (
	"context"

	v1 "github.com/Neakxs/protocel-example/proto/library/v1"
)

type AuthorsRepository interface {
	GetAuthor(ctx context.Context, name string) (*v1.Author, error)
	SaveAuthor(ctx context.Context, author *v1.Author) error
	QueryAuthors(ctx context.Context, pageSize int, pageOffset int) ([]*v1.Author, error)
	DeleteAuthor(ctx context.Context, name string) (*v1.Author, error)
}

type BooksRepository interface {
	GetBook(ctx context.Context, name string) (*v1.Book, error)
	SaveBook(ctx context.Context, book *v1.Book) error
	QueryBooks(ctx context.Context, pageSize int, pageOffset int) ([]*v1.Book, error)
	DeleteBook(ctx context.Context, name string) (*v1.Book, error)
}
