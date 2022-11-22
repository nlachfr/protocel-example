package sqlite

import (
	"context"
	"database/sql"
	"time"

	v1 "github.com/Neakxs/protocel-example/proto/library/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var booksTable = `
	CREATE TABLE IF NOT EXISTS books (
		name varchar(256) PRIMARY KEY,
		author varchar(256) REFERENCES authors(name) NOT NULL,
		publish_date timestamp NOT NULL
	);
`

type booksRepository struct {
	db *sql.DB
}

func (r *booksRepository) GetBook(ctx context.Context, name string) (*v1.Book, error) {
	var publish_date time.Time
	row := r.db.QueryRowContext(ctx, `SELECT * FROM books WHERE name = $1`, name)
	b := &v1.Book{}
	if err := row.Scan(&b.Name, &b.Author, &publish_date); err != nil {
		return nil, err
	}
	b.PublishDate = timestamppb.New(publish_date)
	return b, nil
}

func (r *booksRepository) SaveBook(ctx context.Context, book *v1.Book) error {
	if _, err := r.db.ExecContext(ctx, `
		INSERT INTO authors
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (name)
		DO UPDATE
		SET author = EXCLUDED.author,
			publish_date = EXCLUDED.publish_date
	`, book.Name, book.Author, book.PublishDate.AsTime()); err != nil {
		return err
	}
	return nil
}

func (r *booksRepository) QueryBooks(ctx context.Context, pageSize int, pageOffset int) ([]*v1.Book, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT *
		FROM books
		LIMIT $1
		OFFSET $2
	`, pageSize, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []*v1.Book{}
	for rows.Next() {
		var publish_date time.Time
		b := &v1.Book{}
		if err := rows.Scan(&b.Name, &b.Author, &publish_date); err != nil {
			return nil, err
		}
		b.PublishDate = timestamppb.New(publish_date)
		res = append(res, b)
	}
	return res, nil
}

func (r *booksRepository) DeleteBook(ctx context.Context, name string) (*v1.Book, error) {
	var publish_date time.Time
	row := r.db.QueryRowContext(ctx, `DELETE FROM books WHERE name = $1 RETURNING *`, name)
	b := &v1.Book{}
	if err := row.Scan(&b.Name, &b.Author, &publish_date); err != nil {
		return nil, err
	}
	b.PublishDate = timestamppb.New(publish_date)
	return b, nil
}
