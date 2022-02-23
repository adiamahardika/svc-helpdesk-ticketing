package entity

import "time"

type Category struct {
	Id                int       `json:"id" gorm:"primaryKey"`
	CodeLevel         string    `json:"codeLevel"`
	Name              string    `json:"name"`
	Parent            string    `json:"parent"`
	AdditionalInput_1 string    `json:"additionalInput1"`
	AdditionalInput_2 string    `json:"additionalInput2"`
	AdditionalInput_3 string    `json:"additionalInput3"`
	IsActive          string    `json:"isActive"`
	UpdateAt          time.Time `json:"updateAt"`
}
