package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

// DBLogger struct
type DBLogger struct{}

// BeforeQuery before query string
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

// AfterQuery after query string
func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// New connect pg db
func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
