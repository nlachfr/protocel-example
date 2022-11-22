package db

type WorkUnit interface {
	Authors() AuthorsRepository
	Books() BooksRepository
}
