package service

import (
	"github.com/notesite/models"
	"github.com/notesite/course"
)

type CourseServiceImpl struct {
	courseRepo course.CourseRepository
}

// NewCourseServiceImpl  returns new CourseServiceImpl
func NewCourseServiceImpl(CourseRepo course.CourseRepository) course.CourseService {
	return &CourseServiceImpl{courseRepo: CourseRepo}
}
// Courses returns the courses
func (crs *CourseServiceImpl) Courses() ([]models.Course, []error) {

	crs, errs := crs.courseRepo.Courses()
	if len(errs) > 0 {
		return nil, errs
	}
	return crs, errs

}

// Course retrievs a given course course by its id
func (crs *CourseServiceImpl) Course(id uint) (*models.Course, []error) {
	cr, errs := crs.courseRepo.Course(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cr, errs

}

// UpdateCourse updates a given course course
func (crs *CourseServiceImpl) UpdateCourse(course *models.Course) (*models.Course, []error) {
	cr, errs := crs.courseRepo.UpdateCourse(course)
	if len(errs) > 0 {
		return nil, errs
	}
	return cr, errs

}

// DeleteCourse deletes a given course course
func (crs *CourseServiceImpl) DeleteCourse(id uint) (*models.Course, []error) {

	cr, errs := crs.courseRepo.DeleteCourse(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cr, errs
}

// StoreCourse stores a given course course
func (crs *CourseServiceImpl) StoreCourse(course *models.Course) (*models.Course, []error) {

	cr, errs := crs.courseRepo.StoreCourse(course)
	if len(errs) > 0 {
		return nil, errs
	}
	return cr, errs
}
