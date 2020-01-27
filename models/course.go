package models

type Course struct{
    CourseId uint
    CourseName string `gorm:"type:varchar(255);not null"`
    Users []User
}