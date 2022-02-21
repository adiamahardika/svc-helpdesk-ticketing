package entity

import "time"

type LgServiceActivities struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	LogId          string    `json:"log_id"`
	RequestFrom    string    `json:"request_from"`
	RequestTo      string    `json:"request_to"`
	RequestData    string    `json:"request_data"`
	ResponseData   string    `json:"response_data"`
	RequestTime    time.Time `json:"request_time"`
	ResponseTime   time.Time `json:"response_time"`
	TotalTime      int       `json:"total_time"`
	HttpStatusCode int       `json:"string"`
	LogDate        time.Time `json:"log_date"`
	LogBy          string    `json:"log_by"`
}
