package taskdto

type BaseTaskDto struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Title string `json:"title"` 
    Description string `json:"description"`
}
