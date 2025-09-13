package service

import (
	"proyecto/internal/dto"
	"proyecto/internal/store"
)

type UserService struct {
	store store.UserStore
}

func NewUserService(s store.UserStore) *UserService {
	return &UserService{store: s}
}

func (s *UserService) GetAllUsers() ([]*dto.UsersQuery, error) {
	users, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
