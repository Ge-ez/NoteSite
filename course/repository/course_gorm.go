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
func (courseRepo *CourseRepositoryImpl) Course(id uint) (*models.Course, []error) {
	course := models.Course{}
	errs := courseRepo.conn.First(&course, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &course, errs
}

// CourseByName retrieves a Course by its name from the database

func (courseRepo *CourseRepositoryImpl) CourseByName(name string) (*models.course, []error) {
	course := models.Course{}
	errs := courseRepo.conn.Find(&course, "name=?", name).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &course, errs
}

// Updatecourse updates a given course course in the database

func (courseRepo *CourseRepositoryImpl) UpdateCourse(course *models.Course) (*models.Course, []error) {
	c := course
	errs := courseRepo.conn.Save(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

// DeleteCourse deletes a given course Course from the database
