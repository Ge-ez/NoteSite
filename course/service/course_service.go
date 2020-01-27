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
