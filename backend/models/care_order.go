package models

type CareOrder struct {
    Task
    Patient uint `json:"patient"`
}
