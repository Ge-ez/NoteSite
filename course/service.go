package course

import (
    "github.com/notesite/models"
)

type CourseService interface{
    Courses() ([]models.Course, []error)
   	Course(id uint) (*models.Course, []error)
   	UpdateCourse(course *models.Course) (*models.Course, []error)
   	DeleteCourse(id uint) (*models.Course, []error)
   	StoreCourse(course *models.Course) (*models.Course, []error)
}
