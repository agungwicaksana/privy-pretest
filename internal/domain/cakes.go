package domain

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type CakeEntity struct {
	ID          uint64    `db:"id" json:"id"`
	Code        *string   `db:"title" json:"title"`
	Description *string   `db:"description" json:"description"`
	Rating      *float32  `db:"rating" json:"rating"`
	Image       *string   `db:"image" json:"image"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CakeRepository interface {
	BeginTrx(ctx context.Context) (*sql.Tx, error)
	Find(ctx context.Context, start, limit int, sortBy string) (data []CakeEntity, err error)
	FindOne(ctx context.Context, id string) (data []CakeEntity, err error)
	// Insert(ctx context.Context, data CakeEntity) (result sql.Result, err error)
	// Update(ctx context.Context, data CakeEntity) (result sql.Result, err error)
	// Delete(ctx context.Context, data CakeEntity) (result sql.Result, err error)
}

type repositoryCake struct {
	db *sqlx.DB
}

func NewCakeRepo(db *sqlx.DB) *repositoryCake {
	if db == nil {
		panic("client is nil")
	}

	return &repositoryCake{db}
}

func (r *repositoryCake) BeginTrx(ctx context.Context) (*sql.Tx, error) {
	opts := sql.TxOptions{
		Isolation: 1,
		ReadOnly:  false,
	}
	tx, err := r.db.BeginTx(ctx, &opts)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r *repositoryCake) Find(ctx context.Context, start, limit int, orderBy string) (data []CakeEntity, err error) {
	data = make([]CakeEntity, 0)
	query := fmt.Sprintf("SELECT * FROM cakes ORDER BY %s LIMIT ?, ?", orderBy)
	fmt.Println(orderBy)
	err = r.db.SelectContext(ctx, &data, query, start, limit)
	return
}

func (r *repositoryCake) FindOne(ctx context.Context, id string) (data []CakeEntity, err error) {
	err = r.db.SelectContext(ctx, &data, "SELECT * FROM cakes WHERE id = ?", id)
	return
}
