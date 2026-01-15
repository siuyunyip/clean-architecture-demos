package interactor

type Requester interface {
	Generate(request FinancialReportRequest) (FinancialReportResponse, error)
}
