package sort

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
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

// UseSelectBuilder adds sort to squirrel.SelectBuilder
func (s Sort) UseSelectBuilder(b sq.SelectBuilder) sq.SelectBuilder {
	return b.OrderBy(s.column + " " + s.order)
}
