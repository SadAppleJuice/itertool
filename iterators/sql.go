package iterators

import (
	"database/sql"

	"github.com/SadAppleJuice/itertool"
)

type RowsIterator struct {
	rows *sql.Rows
}

func NewRowsIterator(rows *sql.Rows) *RowsIterator {
	return &RowsIterator{
		rows: rows,
	}
}

func (s *RowsIterator) Next() *itertool.Optional[*sql.Rows] {
	if !s.rows.Next() {
		s.rows.Close()
		return itertool.Null[*sql.Rows]()
	}
	return itertool.Of(s.rows)
}
