package repository

import (
    "github.com/notesite/models"
    "github.com/notesite/upload"
    "github.com/jinzhu/gorm"
)

type UploadRepositoryImpl struct{
    conn gorm.DB
}
