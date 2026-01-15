package interactor

import "siuyunyip.io/ocp-demo/entity"

type FinancialReportRequest struct {
	Year int
}

type FinancialReportResponse struct {
	Reports []*entity.FinancialRecord
}
