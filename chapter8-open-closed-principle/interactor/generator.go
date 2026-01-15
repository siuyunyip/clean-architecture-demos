package interactor

import (
	"siuyunyip.io/ocp-demo/entity"
)

type FinancialDataGateway interface {
	GetFinancialRecords(year int) ([]*entity.FinancialRecord, error)
}

type FinancialReportGenerator struct {
	Gateway FinancialDataGateway
}

func NewFinancialReportGenerator(gw FinancialDataGateway) *FinancialReportGenerator {
	return &FinancialReportGenerator{
		Gateway: gw,
	}
}

func (g FinancialReportGenerator) Generate(request FinancialReportRequest) (FinancialReportResponse, error) {
	records, err := g.Gateway.GetFinancialRecords(request.Year)
	if err != nil {
		panic(err)
	}
	reports := make([]*entity.FinancialRecord, len(records))
	for i, record := range records {
		reports[i] = &entity.FinancialRecord{
			Account: record.Account,
			Amount:  record.Amount,
		}
	}

	return FinancialReportResponse{Reports: reports}, nil
}
