package upload

import "github.com/notesite/models"

type UploadService interface{
    StoreNote(notes *models.Upload) (*models.Upload, []error)
     Notes() ([]models.Upload,[]error)
     Note(id uint) (*models.Upload,[]error)
     UpdateNote(note *models.Upload) (*models.Upload,[]error)
     DeleteNote(id uint) (*models.Upload,[]error)
     NotesByStatus(status string) ([]models.Upload,[]error)
     NotesByUploader(id uint) ([]models.Upload,[]error)
     NotesByCourseName(course string) ([]models.Upload,[]error)
 }
