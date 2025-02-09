package models

type Task struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Title string `json:"title"`
    Description string `json:"description"`
    ParentTask uint `json:"parentTask"`
    Blocks uint `json:"Blocks"`
}
