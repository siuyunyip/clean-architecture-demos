package presenter

import (
	"fmt"

	"siuyunyip.io/ocp-demo/interactor"
)

type PrintPresenter struct{}

func (p PrintPresenter) Present(rsp interactor.FinancialReportResponse) error {
	fmt.Println("=== Financial Report ===")
	for _, r := range rsp.Reports {
		fmt.Printf("%s .... %s\n", r.Account, r.Amount)
	}

	return nil
}
