package sort

import (
	"strings"
)

type Sort struct {
	column string
	order  string
}

func New(column, order string) Sort {
	order = strings.ToUpper(order)
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}

	return Sort{
		column: column,
		order:  order,
	}
}

// Expression returns OrderBy expression like `column DESC`.
func (s Sort) Expression() string {
	return s.column + " " + s.order
}
