package controller

import (
	"testing"

	"siuyunyip.io/ocp-demo/database"
	"siuyunyip.io/ocp-demo/interactor"
	"siuyunyip.io/ocp-demo/presenter"
	"siuyunyip.io/ocp-demo/view"
)

func TestFinancialReportController_Handle(t *testing.T) {
	controller := NewController(presenter.NewScreenViewPresenter(view.WebView{}),
		interactor.NewFinancialReportGenerator(&database.FinancialDataMapper{}))

	err := controller.Handle(1)
	if err != nil {
		return
	}
}
