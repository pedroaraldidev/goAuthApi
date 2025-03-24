package services

import (
    "goAuthApi/entities"
    "goAuthApi/repositories"
)

type UserService struct {
    repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Authenticate(username, password string) (*entities.User, error) {
    user, err := s.repo.FindByUsername(username)
    if err != nil {
        return nil, err
    }
    if user == nil || user.Password != password {
        return nil, nil
    }
    return user, nil
}