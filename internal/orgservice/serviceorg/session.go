package serviceorg

import (
	"database/sql"
	"errors"
	"proyecto/internal/orgservice/dtoorg"

	"github.com/google/uuid"
)

// Find all organizations where the user is a member
type StatusGettingAllUserMemberships int

const (
	UserMembershipsFound StatusGettingAllUserMemberships = iota
	NoUserMembershipsFound
	ErrorGettingUserMemberships
)

func (s *ServiceOrg) GetAllUserMemberships(userID uuid.UUID) ([]dtoorg.UserMembershipsQuery, StatusGettingAllUserMemberships, error) {

	memberships, err := s.store.GetUserMemberships(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NoUserMembershipsFound, nil

		}
		return nil, ErrorGettingUserMemberships, err
	}

	return memberships, UserMembershipsFound, nil
}

// Check if a user is a active member of an organization and return the membership ID
type StatusCheckingUserMembership int

const (
	NotFoundMembership StatusCheckingUserMembership = iota
	FoundMembership
	ErrorCheckingMembership
)

func (s *ServiceOrg) CheckUserMembership(userID, memID uuid.UUID) (uuid.UUID, StatusCheckingUserMembership, error) {

	memID, err := s.store.CheckUserMembership(userID, memID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, NotFoundMembership, nil
		}
		return uuid.Nil, ErrorCheckingMembership, err
	}

	return memID, FoundMembership, nil
}
