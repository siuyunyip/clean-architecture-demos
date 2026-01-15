package controller

import "siuyunyip.io/ocp-demo/interactor"

type FinancialReportController struct {
	presenter ReportPresenter
	requester interactor.Requester
}

func NewController(p ReportPresenter, r interactor.Requester) *FinancialReportController {
	return &FinancialReportController{
		presenter: p,
		requester: r,
	}
}

func (c *FinancialReportController) Handle(year int) error {
	req := interactor.FinancialReportRequest{Year: year}
	rsp, err := c.requester.Generate(req)
	if err != nil {
		panic(err)
	}

	if err = c.presenter.Present(rsp); err != nil {
		panic(err)
	}

	return nil
}
