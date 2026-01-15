package controller

import "siuyunyip.io/ocp-demo/interactor"

type ReportPresenter interface {
	Present(rsp interactor.FinancialReportResponse) error
}
