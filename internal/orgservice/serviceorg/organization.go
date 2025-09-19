package serviceorg

import (
	"proyecto/internal/orgservice/dtoorg"
	"proyecto/internal/orgservice/modelorg"
	"proyecto/internal/shared/utils"

	"github.com/google/uuid"
)

func (s *ServiceOrg) GetAllOrganizationsByUserID(userID uuid.UUID) ([]dtoorg.CreatedOrganization, error) {
	return s.store.GetAllOrganizationsByUserID(userID.String())
}

type StatusCreatingPersonalOrganization int

const (
	PersonalOrganizationCreated StatusCreatingPersonalOrganization = iota
	PersonalOrganizationAlreadyExists
)

func (s *ServiceOrg) CreateNewPersonalOrganization(userId uuid.UUID, orgName string) (StatusCreatingPersonalOrganization, dtoorg.CreatedOrganization, error) {
	var createdOrg dtoorg.CreatedOrganization

	// Check if personal organization already exists for the user.
	// Only one personal organization is allowed per user.
	exist, err := s.store.ExistOrganizationPersonal(userId.String())
	if err != nil {
		return 0, createdOrg, err
	}
	if exist {
		return PersonalOrganizationAlreadyExists, createdOrg, nil
	}

	// Create new personal organization
	newOrg := modelorg.Organizations{
		ID:        uuid.New(),
		OwnerID:   userId,
		Name:      orgName,
		Type:      "personal",
		CreatedAt: utils.TimeNow(),
	}

	createdOrg, err = s.store.CreateNewOrganization(&newOrg)
	if err != nil {
		return 0, createdOrg, err
	}

	return PersonalOrganizationCreated, createdOrg, nil
}

type statusCreatingCompanyOrganization int

const (
	CompanyOrganizationCreated statusCreatingCompanyOrganization = iota
	CompanyOrganizationNameExists
	MaxCompanyOrganizationsReached
)

func (s *ServiceOrg) CreateNewCompanyOrganization(userId uuid.UUID, orgName string) (statusCreatingCompanyOrganization, dtoorg.CreatedOrganization, error) {
	var createdOrg dtoorg.CreatedOrganization
	// Check if company organization already exists for the user.
	// Multiple company organizations are allowed per user.

	companies, err := s.store.GetAllOrganizationsByUserID(userId.String())
	if err != nil {
		return 0, createdOrg, err
	}

	// Limit to a maximum of 3 company organizations
	if len(companies) >= 3 {
		return MaxCompanyOrganizationsReached, createdOrg, nil
	}

	// Check if organization name already exists
	for _, org := range companies {
		if org.Name == orgName {
			return CompanyOrganizationNameExists, createdOrg, nil
		}
	}

	// Create new company organization
	newOrg := modelorg.Organizations{
		ID:        uuid.New(),
		OwnerID:   userId,
		Name:      orgName,
		Type:      "company",
		CreatedAt: utils.TimeNow(),
	}

	createdOrg, err = s.store.CreateNewOrganization(&newOrg)
	if err != nil {
		return 0, createdOrg, err
	}

	return CompanyOrganizationCreated, createdOrg, nil
}
