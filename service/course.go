package service

import (
	"github.com/lucasfloriani/go-mongo/model"
)

// courseDAO specifies the interface of the course DAO needed by CourseService.
type courseDAO interface {
	All(offset, limit int) ([]model.Course, error)
	Count() (int, error)
	Get(id string) (*model.Course, error)
	Create(u *model.Course) error
	Update(u *model.Course) error
	Delete(u *model.Course) error
}

// CourseService provides services related with courses.
type CourseService struct {
	dao courseDAO
}

// NewCourseService creates a new CourseService with the given course DAO.
func NewCourseService(dao courseDAO) *CourseService {
	return &CourseService{dao}
}

// Count returns the number of courses.
func (s *CourseService) Count() (int, error) {
	return s.dao.Count()
}

// Query returns the courses with the specified offset and limit.
func (s *CourseService) Query(offset, limit int) ([]model.Course, error) {
	return s.dao.All(offset, limit)
}

// Get returns the course with the specified the course ID.
func (s *CourseService) Get(id string) (*model.Course, error) {
	return s.dao.Get(id)
}

// Create creates a new course.
func (s *CourseService) Create(u *model.Course) (*model.Course, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

// Update updates the course with the specified ID.
func (s *CourseService) Update(u *model.Course) (*model.Course, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(u); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete deletes the course with the specified ID.
func (s *CourseService) Delete(id string) (*model.Course, error) {
	course, err := s.dao.Get(id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(course)
	return course, err
}
