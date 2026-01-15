package database

import "siuyunyip.io/ocp-demo/entity"

type FinancialDataMapper struct{}

func (m FinancialDataMapper) GetFinancialRecords(year int) ([]*entity.FinancialRecord, error) {
	return Records[year], nil
}
