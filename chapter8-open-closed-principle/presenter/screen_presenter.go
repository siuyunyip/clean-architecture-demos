package presenter

import (
	"fmt"

	"siuyunyip.io/ocp-demo/interactor"
)

type ScreenViewModel struct {
	Lines []string
}

type ScreenViewPresenter struct {
	ViewModel      *ScreenViewModel
	ScreenRenderer ScreenRenderer
}

func NewScreenViewPresenter(renderer ScreenRenderer) *ScreenViewPresenter {
	return &ScreenViewPresenter{
		ViewModel:      &ScreenViewModel{},
		ScreenRenderer: renderer,
	}
}

func (p ScreenViewPresenter) Present(rsp interactor.FinancialReportResponse) error {
	p.ViewModel.Lines = nil
	for _, r := range rsp.Reports {
		p.ViewModel.Lines = append(
			p.ViewModel.Lines,
			fmt.Sprintf("%s: %s", r.Account, r.Amount),
		)
	}

	if err := p.ScreenRenderer.Render(p.ViewModel); err != nil {
		panic(err)
	}

	return nil
}
