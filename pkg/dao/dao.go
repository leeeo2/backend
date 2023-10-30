package dao

import "gorm.io/gorm"

type DescribeInput struct {
	Query      string
	Params     []interface{}
	PageSize   int
	PageNumber int
	Order      string

	NextToken           string
	MaxResults          int
	NextTOkenPagination bool
}

func GetDescribeTx(db *gorm.DB, input *DescribeInput) *gorm.DB {
	tx := db
	tx = tx.Where(input.Query, input.Params...)

	if input.PageSize > 0 && input.PageNumber > 0 {
		tx = tx.Limit(input.PageSize).Offset((input.PageNumber - 1) * input.PageSize)
	}

	if len(input.Order) != 0 {
		tx = tx.Order(input.Order)
	}
	return tx
}
