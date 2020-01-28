package service

import(
    "github.com/notesite/models"
    "github.com/notesite/upload"
)

type UploadServiceImpl struct{
    uploadrepo upload.UploadRepository
}

func NewUploadSericeImpl(upldServ *UploadServiceImpl) *UploadServiceImpl{
    return &UploadServiceImpl{uploadrepo:upldServ}
}

func (usi *UploadServiceImpl) StoreNote(notes *models.Upload) (*models.Upload, []error) {
	note,err := usi.uploadrepo.StoreNote(notes)
    	if len(err) > 0 {
    		return note,err
    	}
    	return note,nil
}
func (usi *UploadServiceImpl) Notes() (*models.Upload,[]error){
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

func (usi *UploadServiceImpl) NoteByStatus(status string) (*models.Upload,[]error){
        note, err := usi.uploadrepo.NoteByStatus(status)
          	if len(err) > 0 {
          		return nil, err
          	}
          	return note, err
}

func (usi *UploadServiceImpl) NoteByUploader(id uint) (*models.Upload,[]error){
        note, err := usi.uploadrepo.NoteByUploader(id)
          	if len(err) > 0 {
          		return nil, err
          	}
          	return note, err
}

func (usi *UploadServiceImpl) NoteByCourseName(course string) (*models.Upload,[]error){
        note, err := usi.uploadrepo.NoteByCourseName(course)
          	if len(err) > 0 {
          		return nil, err
          	}
          	return note, err
}
