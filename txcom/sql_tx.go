package txcom

import (
	"database/sql"

	"github.com/w-woong/common/port"
)

type SqlTxBeginner struct {
	db *sql.DB
}

func NewSqlTxBeginner(db *sql.DB) *SqlTxBeginner {
	return &SqlTxBeginner{db}
}

func (t *SqlTxBeginner) Begin() (port.TxController, error) {
	return t.db.Begin()
}

// Begin starts transaction returning port.TxController that commits or rollbacks
func (t *SqlTxBeginner) BeginR() (port.TxController, error) {
	return t.Begin()
}

type NopTxBeginner struct{}

func NewNopTxBeginner() *NopTxBeginner {
	return &NopTxBeginner{}
}
func (o *NopTxBeginner) Begin() (port.TxController, error) {
	return NewNopTxController(), nil
}

type NopTxController struct{}

func NewNopTxController() *NopTxController {
	return &NopTxController{}
}
func (o *NopTxController) Commit() error {
	return nil
}
func (o *NopTxController) Rollback() error {
	return nil
}
