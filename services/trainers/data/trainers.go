package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v9"
)

type Trainer struct {
	ID        int64
	Name      string
	Email     string
	Address   string
	Phone     string
	NhsNumber string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TrainersModel struct {
	db    *sql.DB
	redis *redis.Client
}

type TrainersModeler interface {
	Insert(context.Context, *Trainer) error
	Update(parent context.Context, id int64, t *Trainer) error
}

func NewTrainersModel(db *sql.DB, redis *redis.Client) TrainersModel {
	return TrainersModel{
		db:    db,
		redis: redis,
	}
}

func (m TrainersModel) Insert(parent context.Context, t *Trainer) error {
	query := sq.
		Insert("trainers").
		Columns("name", "email", "address", "nhs_number", "phone").
		Values(t.Name, t.Email, t.Address, t.NhsNumber, t.Phone).
		Suffix("RETURNING id, created_at, updated_at").
		RunWith(m.db).
		PlaceholderFormat(sq.Dollar)

	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()

	return query.QueryRowContext(ctx).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (m TrainersModel) Update(parent context.Context, id int64, t *Trainer) error {
	query := sq.
		Update("trainers").
		SetMap(map[string]interface{}{
			"name":       t.Name,
			"email":      t.Email,
			"address":    t.Address,
			"nhs_number": t.NhsNumber,
			"phone":      t.Phone,
		}).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING updated_at").
		RunWith(m.db).
		PlaceholderFormat(sq.Dollar)

	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()

	err := query.QueryRowContext(ctx).Scan(&t.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return errors.New("edit conflict")
		default:
			return err
		}
	}

	return nil
}
