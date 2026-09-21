package repos

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"school-portal/internal/domain"
)

type TeacherRepo interface {
	// FindByIDNumberAndEmployeeNumber matches both values against the same record in a single
	// query, so a caller can never learn from the result which of the two fields was wrong.
	FindByIDNumberAndEmployeeNumber(ctx context.Context, idNumber string, employeeNumber string) (*domain.Teacher, error)
	UpdateAccountPasswordHash(ctx context.Context, id string, passwordHash string) error
}

type teacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) TeacherRepo {
	return &teacherRepo{db: db}
}

func (r *teacherRepo) FindByIDNumberAndEmployeeNumber(ctx context.Context, idNumber string, employeeNumber string) (*domain.Teacher, error) {
	const query = `
		SELECT id, id_number, employee_number, full_name, account_username, account_password_hash, created_at, updated_at
		FROM teachers
		WHERE id_number = $1 AND employee_number = $2
	`
	row := r.db.QueryRow(ctx, query, idNumber, employeeNumber)
	return scanTeacher(row)
}

func (r *teacherRepo) UpdateAccountPasswordHash(ctx context.Context, id string, passwordHash string) error {
	const query = `
		UPDATE teachers
		SET account_password_hash = $1, updated_at = now()
		WHERE id = $2
	`
	_, err := r.db.Exec(ctx, query, passwordHash, id)
	return err
}

func scanTeacher(row rowScanner) (*domain.Teacher, error) {
	var t domain.Teacher
	if err := row.Scan(&t.ID, &t.IDNumber, &t.EmployeeNumber, &t.FullName, &t.AccountUsername, &t.AccountPasswordHash, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}
	return &t, nil
}
