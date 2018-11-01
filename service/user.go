package service

import (
	"github.com/lucasfloriani/go-mongo/model"
)

// userDAO specifies the interface of the user DAO needed by UserService.
type userDAO interface {
	All(offset, limit int) ([]model.User, error)
	Count() (int, error)
	Get(id string) (*model.User, error)
	Create(u *model.User) error
	Update(u *model.User) error
	Delete(u *model.User) error
}

// UserService provides services related with users.
type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Count returns the number of users.
func (s *UserService) Count() (int, error) {
	return s.dao.Count()
}

// Query returns the users with the specified offset and limit.
func (s *UserService) Query(offset, limit int) ([]model.User, error) {
	return s.dao.All(offset, limit)
}

// Get returns the user with the specified the user ID.
func (s *UserService) Get(id string) (*model.User, error) {
	return s.dao.Get(id)
}

// Create creates a new user.
func (s *UserService) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

// Update updates the user with the specified ID.
func (s *UserService) Update(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(u); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete deletes the user with the specified ID.
func (s *UserService) Delete(id string) (*model.User, error) {
	user, err := s.dao.Get(id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(user)
	return user, err
}
