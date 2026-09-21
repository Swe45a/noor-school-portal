package web

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	apperrors "school-portal/internal/errors"
	"school-portal/internal/repos"
	"school-portal/internal/utils"
)

// TeacherLoginResult carries the account login information shown to the teacher once their
// ID Number and Employee Number have been verified as belonging to the same record.
type TeacherLoginResult struct {
	FullName        string
	AccountUsername string
	AccountPassword string
}

type TeacherService interface {
	Login(ctx context.Context, idNumber string, employeeNumber string) (*TeacherLoginResult, error)
}

type teacherService struct {
	teachers repos.TeacherRepo
}

func NewTeacherService(teachers repos.TeacherRepo) TeacherService {
	return &teacherService{teachers: teachers}
}

func (s *teacherService) Login(ctx context.Context, idNumber string, employeeNumber string) (*TeacherLoginResult, error) {
	teacher, err := s.teachers.FindByIDNumberAndEmployeeNumber(ctx, idNumber, employeeNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrTeacherNotFound
		}
		return nil, fmt.Errorf("finding teacher: %w", err)
	}

	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return nil, fmt.Errorf("generating password: %w", err)
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	if err := s.teachers.UpdateAccountPasswordHash(ctx, teacher.ID, hash); err != nil {
		return nil, fmt.Errorf("updating account password hash: %w", err)
	}

	return &TeacherLoginResult{
		FullName:        teacher.FullName,
		AccountUsername: teacher.AccountUsername,
		AccountPassword: password,
	}, nil
}
