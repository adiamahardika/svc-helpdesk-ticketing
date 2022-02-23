package entity

import "time"

type LgServiceActivities struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	LogId          string    `json:"logId"`
	RequestFrom    string    `json:"requestFrom"`
	RequestTo      string    `json:"requestTo"`
	RequestData    string    `json:"requestData"`
	ResponseData   string    `json:"responseData"`
	RequestTime    time.Time `json:"requestTime"`
	ResponseTime   time.Time `json:"responseTime"`
	TotalTime      int       `json:"totalTime"`
	HttpStatusCode int       `json:"httpStatusCode"`
	LogDate        time.Time `json:"logDate"`
	LogBy          string    `json:"logBy"`
}
