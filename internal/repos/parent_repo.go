package repos

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"school-portal/internal/domain"
)

type ParentRepo interface {
	FindByParentID(ctx context.Context, parentID string) (*domain.Parent, error)
	UpdateAccountPasswordHash(ctx context.Context, id string, passwordHash string) error
}

type parentRepo struct {
	db *pgxpool.Pool
}

func NewParentRepo(db *pgxpool.Pool) ParentRepo {
	return &parentRepo{db: db}
}

func (r *parentRepo) FindByParentID(ctx context.Context, parentID string) (*domain.Parent, error) {
	const query = `
		SELECT id, parent_id, full_name, account_username, account_password_hash, created_at, updated_at
		FROM parents
		WHERE parent_id = $1
	`
	row := r.db.QueryRow(ctx, query, parentID)
	return scanParent(row)
}

func (r *parentRepo) UpdateAccountPasswordHash(ctx context.Context, id string, passwordHash string) error {
	const query = `
		UPDATE parents
		SET account_password_hash = $1, updated_at = now()
		WHERE id = $2
	`
	_, err := r.db.Exec(ctx, query, passwordHash, id)
	return err
}

// rowScanner is shared by every repo's scan helper in this package.
type rowScanner interface {
	Scan(dest ...interface{}) error
}

func scanParent(row rowScanner) (*domain.Parent, error) {
	var p domain.Parent
	if err := row.Scan(&p.ID, &p.ParentID, &p.FullName, &p.AccountUsername, &p.AccountPasswordHash, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return nil, err
	}
	return &p, nil
}
