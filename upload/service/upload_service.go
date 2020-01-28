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
