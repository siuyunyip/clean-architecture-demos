package database

import "siuyunyip.io/ocp-demo/entity"

var (
	Records = map[int][]*entity.FinancialRecord{
		1: []*entity.FinancialRecord{
			&entity.FinancialRecord{
				Amount:  "2000",
				Account: "1",
			},
			&entity.FinancialRecord{
				Amount:  "1000",
				Account: "2",
			},
		},
		2: []*entity.FinancialRecord{
			&entity.FinancialRecord{
				Amount:  "3000",
				Account: "1",
			},
			&entity.FinancialRecord{
				Amount:  "2000",
				Account: "2",
			},
		},
	}
)
