package sqlite

import (
	"database/sql"

	"github.com/Neakxs/protocel-example/internal/db"
)

type workUnit struct {
	db *sql.DB
}

func (u *workUnit) Authors() db.AuthorsRepository {
	return &authorsRepository{db: u.db}
}
func (u *workUnit) Books() db.BooksRepository {
	return &booksRepository{db: u.db}
}

func NewWorkUnit(db *sql.DB) db.WorkUnit {
	return &workUnit{db: db}
}
