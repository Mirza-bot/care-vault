package userdtos

type BaseUserDto struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    Email string `json:"email" gorm:"unique"`
    IsAccount bool `json:"isAccount"`
}
