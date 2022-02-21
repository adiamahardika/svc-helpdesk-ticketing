package service

import (
	"log"
	"regexp"
	"strconv"
	"svc-myg-ticketing/entity"
	"svc-myg-ticketing/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type LogServiceInterface interface {
	CreateLog(context *gin.Context, request string, response string, start time.Time, httpStatus int)
}

type logService struct {
	repository repository.LogRepositoryInterface
}

func LogService(repository repository.LogRepositoryInterface) *logService {
	return &logService{repository}
}

func (logService *logService) CreateLog(context *gin.Context, request string, response string, start time.Time, httpStatus int) {
	now := time.Now()
	check, err := regexp.Compile("[^a-zA-Z0-9]+")
	var request_by string

	if context.Request.Header.Get("request-by") != "" {
		request_by = context.Request.Header.Get("request-by")
	} else {
		request_by = context.ClientIP()
	}
	if err != nil {
		log.Fatal(err)
	}
	end_time := check.ReplaceAllString(now.Format("150405.000000000"), "")
	conv_end_time, _ := strconv.Atoi(end_time)
	start_time := check.ReplaceAllString(start.Format("150405.000000000"), "")
	conv_start_time, _ := strconv.Atoi(start_time)

	parse_ip := check.ReplaceAllString(context.ClientIP(), "")
	log_request := entity.LgServiceActivities{
		LogId:          parse_ip + start.Format("20060102150405"),
		RequestFrom:    request_by,
		RequestTo:      context.Request.RequestURI,
		RequestData:    request,
		ResponseData:   response,
		RequestTime:    start,
		ResponseTime:   now,
		TotalTime:      conv_end_time - conv_start_time,
		HttpStatusCode: httpStatus,
		LogDate:        now,
		LogBy:          request_by,
	}
	logService.repository.CreateLog(log_request)

}
