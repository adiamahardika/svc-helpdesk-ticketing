package entity

import "time"

type Category struct {
	Id                int       `json:"id" gorm:"primaryKey"`
	CodeLevel         string    `json:"code_level"`
	Name              string    `json:"name"`
	Parent            string    `json:"parent"`
	AdditionalInput_1 string    `json:"additional_input_1"`
	AdditionalInput_2 string    `json:"additional_input_2"`
	AdditionalInput_3 string    `json:"additional_input_3"`
	IsActive          string    `json:"is_active"`
	UpdateAt          time.Time `json:"update_at"`
}
