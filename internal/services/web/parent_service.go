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

// ParentLoginResult carries the account login information shown to the parent once their
// Parent ID has been verified against the database.
type ParentLoginResult struct {
	FullName        string
	AccountUsername string
	AccountPassword string
}

type ParentService interface {
	// Login looks up the parent by their Parent ID and, if found, issues a fresh account
	// password to display. There is no OTP step: possession of a valid Parent ID is the
	// only factor required, matched by the caller's rate-limited access to this endpoint.
	Login(ctx context.Context, parentID string) (*ParentLoginResult, error)
}

type parentService struct {
	parents repos.ParentRepo
}

func NewParentService(parents repos.ParentRepo) ParentService {
	return &parentService{parents: parents}
}

func (s *parentService) Login(ctx context.Context, parentID string) (*ParentLoginResult, error) {
	parent, err := s.parents.FindByParentID(ctx, parentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrParentNotFound
		}
		return nil, fmt.Errorf("finding parent: %w", err)
	}

	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return nil, fmt.Errorf("generating password: %w", err)
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	if err := s.parents.UpdateAccountPasswordHash(ctx, parent.ID, hash); err != nil {
		return nil, fmt.Errorf("updating account password hash: %w", err)
	}

	return &ParentLoginResult{
		FullName:        parent.FullName,
		AccountUsername: parent.AccountUsername,
		AccountPassword: password,
	}, nil
}
