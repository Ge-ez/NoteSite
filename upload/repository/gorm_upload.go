package repository

import (
	"fmt"
	"github.com/notesite/models"
    "github.com/jinzhu/gorm"
	"github.com/techreview/entities"
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

func (upld *UploadRepositoryImpl) RateNote(rating *models.Rating) {

	errs := upld.conn.Where("article_id = ? AND user_id= ?", rating.FileID, rating.UserID).First(&rating).GetErrors()

	if len(errs) > 0 {

		errs := upld.conn.Create(rating).GetErrors()
		if len(errs) > 0 {
			fmt.Println("error storing  ratings")
			return
		}
		fmt.Println("stored   ratings")
		fmt.Println("query worked")
		return

	} else{
		errs := upld.conn.Delete(rating).GetErrors()
		if len(errs) > 0 {
			fmt.Println("error while deleting  ratings", errs)
			return
		}
		fmt.Println("deleted answer  ratings successfuly")
		return
	}
}

func (upld *UploadRepositoryImpl)  NoteRateCount(FileId uint) int{
	var count int
	var noteRating models.Rating
	//questionFollows := entities.QuestionFollow{}
	upld.conn.Model(&noteRating).Where("article_id = ?", FileId).Count(&count)
	// update the count value in articles table
	upld.conn.Model(&entities.Article{}).Where("id = ?", FileId).UpdateColumn("number_of_ratings", count)
	// update the average ratings in articles table
	upld.conn.Model(&entities.Article{}).Where("id = ?", FileId).UpdateColumn("average_rating", count/5) // it is doing it just the data

	fmt.Println("counted succesfully, with value: ", count)

	return count
}
