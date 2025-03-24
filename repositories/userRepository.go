package repositories

import "goAuthApi/entities"

type UserRepository interface {
    FindByUsername(username string) (*entities.User, error)
}

type InMemoryUserRepository struct {
    users []entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users: []entities.User{
            {ID: 1, Username: "user1", Password: "password1"},
        },
    }
}

func (r *InMemoryUserRepository) FindByUsername(username string) (*entities.User, error) {
    for _, user := range r.users {
        if user.Username == username {
            return &user, nil
        }
    }
    return nil, nil
}