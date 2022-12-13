package db

import "context"

type WorkUnit interface {
	Authors() AuthorsRepository
	Books() BooksRepository
	Migrate(ctx context.Context) error
}
