package service

import(
    "github.com/notesite/models"
    "github.com/notesite/upload"
)

type UploadServiceImpl struct{
    uploadrepo upload.UploadRepository
}

func NewUploadServiceImpl(upldServ upload.UploadRepository) upload.UploadService{
    return &UploadServiceImpl{uploadrepo:upldServ}
}

func (usi *UploadServiceImpl) StoreNote(notes *models.Upload) (*models.Upload, []error) {
	note,err := usi.uploadrepo.StoreNote(notes)
    	if len(err) > 0 {
    		return note,err
    	}
    	return note,nil
}
func (usi *UploadServiceImpl) Notes() ([]models.Upload,[]error){
        notes, err := usi.uploadrepo.Notes()
        	if len(err) > 0 {
        		return nil, err
        	}
        	return notes, err
}

func (usi *UploadServiceImpl) Note(id uint) (*models.Upload,[]error){
    note, err := usi.uploadrepo.Note(id)
    	if len(err) > 0 {
    		return nil, err
    	}
    	return note, nil
}

func (usi *UploadServiceImpl) UpdateNote(update *models.Upload) (*models.Upload,[]error){
   note,err := usi.uploadrepo.UpdateNote(update)
   	if len(err) > 0 {
   		return note,err
   	}
   	return note,nil
}
func (usi *UploadServiceImpl) DeleteNote(id uint) (*models.Upload,[]error){
    note,err := usi.uploadrepo.DeleteNote(id)
    	if len(err) > 0 {
    		return nil,err
    	}
    	return note,nil
}


func (usi *UploadServiceImpl) NotesByStatus(status string) ([]models.Upload,[]error){
        note, err := usi.uploadrepo.NotesByStatus(status)
          	if len(err) > 0 {
          		return nil, err
          	}
          	return note, err
}

func (usi *UploadServiceImpl) NotesByUploader(id uint) ([]models.Upload,[]error){
        note, err := usi.uploadrepo.NotesByUploader(id)
          	if len(err) > 0 {
          		return nil, err
          	}
          	return note, err
}

func (usi *UploadServiceImpl) NotesByCourseName(course string) ([]models.Upload,[]error){
        note, err := usi.uploadrepo.NotesByCourseName(course)
          	if len(err) > 0 {
          		return nil, err
          	}
          	return note, err
}

