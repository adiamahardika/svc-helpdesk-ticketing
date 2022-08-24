package service

import (
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/repository"
)

type ReportServiceInterface interface {
	GetReport(request *model.GetReportRequest) ([]model.ReportResponse, error)
}

type reportService struct {
	reportRepository repository.ReportRepositoryInterface
}

func ReportService(reportRepository repository.ReportRepositoryInterface) *reportService {
	return &reportService{reportRepository}
}

func (reportService *reportService) GetReport(request *model.GetReportRequest) ([]model.ReportResponse, error) {

	request.EndDate = request.EndDate + " 23:59:59"
	ticket, error := reportService.reportRepository.GetReport(request)

	return ticket, error
}
