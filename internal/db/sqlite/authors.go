package sqlite

import (
	"context"
	"database/sql"
	"time"

	v1 "github.com/Neakxs/protocel-example/proto/library/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var authorsTable = `
	CREATE TABLE IF NOT EXISTS authors (
		name varchar(256) PRIMARY KEY,
		display_name varchar(256) NOT NULL,
		birth_date timestamp NOT NULL,
		death_date timestamp
	);
`

type authorsRepository struct {
	db *sql.DB
}

func (r *authorsRepository) GetAuthor(ctx context.Context, name string) (*v1.Author, error) {
	var (
		birth, death *time.Time
	)
	row := r.db.QueryRowContext(ctx, `SELECT * FROM authors WHERE name = $1`, name)
	a := &v1.Author{}
	if err := row.Scan(&a.Name, &a.DisplayName, &birth, &death); err != nil {
		return nil, err
	}
	a.BirthDate = timestamppb.New(*birth)
	if death != nil {
		a.DeathDate = timestamppb.New(*death)
	}
	return a, nil
}

func (r *authorsRepository) SaveAuthor(ctx context.Context, author *v1.Author) error {
	var death *time.Time
	if author.DeathDate != nil {
		t := author.DeathDate.AsTime()
		death = &t
	}
	if _, err := r.db.ExecContext(ctx, `
		INSERT INTO authors
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (name)
		DO UPDATE
		SET display_name = EXCLUDED.display_name,
			birth_date = EXCLUDED.birth_date,
			death_date = EXCLUDED.death_date
	`, author.Name, author.DisplayName, author.BirthDate.AsTime(), death); err != nil {
		return err
	}
	return nil
}

func (r *authorsRepository) QueryAuthors(ctx context.Context, pageSize int, pageOffset int) ([]*v1.Author, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT *
		FROM authors
		LIMIT $1
		OFFSET $2
	`, pageSize, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []*v1.Author{}
	for rows.Next() {
		var (
			birth, death *time.Time
		)
		a := &v1.Author{}
		if err := rows.Scan(&a.Name, &a.DisplayName, &birth, &death); err != nil {
			return nil, err
		}
		a.BirthDate = timestamppb.New(*birth)
		if death != nil {
			a.DeathDate = timestamppb.New(*death)
		}
		res = append(res, a)
	}
	return res, nil
}

func (r *authorsRepository) DeleteAuthor(ctx context.Context, name string) (*v1.Author, error) {
	var (
		birth, death *time.Time
	)
	row := r.db.QueryRowContext(ctx, `DELETE FROM authors WHERE name = $1 RETURNING *`, name)
	a := &v1.Author{}
	if err := row.Scan(&a.Name, &a.DisplayName, &birth, &death); err != nil {
		return nil, err
	}
	a.BirthDate = timestamppb.New(*birth)
	if death != nil {
		a.DeathDate = timestamppb.New(*death)
	}
	return a, nil
}
