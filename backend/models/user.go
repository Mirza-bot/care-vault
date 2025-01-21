package models

import "time"

type User struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    Email string `json:"email" gorm:"unique"`
    Password string `json:"-"`
    Created time.Time `json:"created" gorm:"autoCreateTime"`
    Role `json:"role"`
}
