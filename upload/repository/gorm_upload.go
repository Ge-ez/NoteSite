package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/notesite/models"
)

type UploadRepositoryImpl struct{
    conn *gorm.DB
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
func (upld *UploadRepositoryImpl) Notes() ([]models.Upload,[]error){
        notes := []models.Upload{}
    	errs := upld.conn.Find(&notes).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return notes, errs
}

func (upld *UploadRepositoryImpl) Note(id uint) (*models.Upload,[]error){
    note := models.Upload{}
    	errs := upld.conn.First(&note, id).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return &note, errs
}

func (upld *UploadRepositoryImpl) UpdateNote(update *models.Upload) (*models.Upload,[]error){
    note := update
    	errs := upld.conn.Save(note).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}
func (upld *UploadRepositoryImpl) DeleteNote(id uint) (*models.Upload,[]error){
    note, errs := upld.Note(id)
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	errs = upld.conn.Delete(note, id).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}


func (upld *UploadRepositoryImpl) NotesByStatus(status string) ([]models.Upload,[]error){
        note := []models.Upload{}
    	errs := upld.conn.Find(&note, "file_status=?", status).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}

func (upld *UploadRepositoryImpl) NotesByUploader(id uint) ([]models.Upload,[]error){
        note := []models.Upload{}
    	errs := upld.conn.Find(&note, "file_uploader=?", id).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}

func (upld *UploadRepositoryImpl) NotesByCourseName(course string) ([]models.Upload,[]error){
        note := []models.Upload{}
    	errs := upld.conn.Find(&note, "file_uploaded_to=?", course).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}
