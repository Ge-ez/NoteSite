package repository

import (
    "github.com/notesite/models"
    "github.com/notesite/upload"
    "github.com/jinzhu/gorm"
)

type UploadRepositoryImpl struct{
    conn gorm.DB
}

func NewUploadRepositoryImpl(db *gorm.DB) *UploadRepositoryImpl{
    return &UploadRepositoryImpl{conn:db}
}
func (upld *UploadRepositoryImpl) StoreNote(notes *models.Upload) (*models.Upload, []error) {
	note := notes
	errs := upld.conn.Create(note).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return note, errs
}
