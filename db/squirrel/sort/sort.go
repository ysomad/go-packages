package sort

import (
	"errors"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

var errEmptyColumn = errors.New("column cannot be empty string")

type Sort struct {
	column string
	order  string
}

func New(column, order string) (Sort, error) {
	if column == "" {
		return Sort{}, errEmptyColumn
	}

	order = strings.ToUpper(order)
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}

	return Sort{
		column: column,
		order:  order,
	}, nil
}

// UseSelectBuilder adds sort to squirrel.SelectBuilder
func (s Sort) UseSelectBuilder(b sq.SelectBuilder) sq.SelectBuilder {
	return b.OrderBy(s.column + " " + s.order)
}
