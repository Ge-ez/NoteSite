package models

import "time"

type Upload struct{
    ID uint
    File_Name string `gorm:"type:varchar(255);not null"`
    File_Description string `gorm:"type:varchar(255);not null"`
    File_Type string `gorm:"type:varchar(255);not null"`
    File_Uploader uint
    File_Uploaded_To string `gorm:"type:varchar(255);not null"`
    File_Uploaded_On time.Time
    File string `gorm:"type:varchar(255);not null"`
    File_Status string `gorm:"type:varchar(255);not null"`
}