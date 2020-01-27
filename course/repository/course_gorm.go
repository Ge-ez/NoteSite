package repository

import (
	"github.com/notesite/models"
	"github.com/notesite/course"
	"github.com/jinzhu/gorm"
)

// CourseRepositoryImpl implements the menu.CourseRepository interface
type CourseRepositoryImpl struct {
	conn *gorm.DB
}

// NewCourseRepositoryImpl returns a new a new object of CourseRepositoryImpl
func NewCourseRepositoryImpl(db *gorm.DB) course.CourseRepository {
	return &CourseRepositoryImpl{conn: db}
}

// Courses returns all course Courses stored in the database
func (courseRepo *CourseRepositoryImpl) Courses() ([]models.Course, []error) {
	courses := []models.Course{}
	errs := courseRepo.conn.Find(&courses).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return courses, errs
}

// course retrieves a course by its id from the database
