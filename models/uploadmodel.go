package models

import "time"

type Upload struct{
    ID uint
    File_Name string
    File_Description string
    File_Type string
    File_Uploader uint
    File_Uploaded_To string
    File_Uploaded_On time.Time
    File string
    File_Status string
}
