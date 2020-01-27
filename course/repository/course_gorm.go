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
