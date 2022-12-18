package txcom

import (
	"database/sql"

	"github.com/w-woong/common"
)

type SqlIsolationLevelSetter struct {
}

func NewSqlIsolationLevelSetter() *SqlIsolationLevelSetter {
	return &SqlIsolationLevelSetter{}
}

func (a *SqlIsolationLevelSetter) SetReadUncommitted(tx common.TxController) error {
	_, err := tx.(*sql.Tx).Exec("set transaction isolation level read uncommitted")
	if err != nil {
		return err
	}

	return nil
}

func (a *SqlIsolationLevelSetter) SetReadCommitted(tx common.TxController) error {
	_, err := tx.(*sql.Tx).Exec("set transaction isolation level read committed")
	if err != nil {
		return err
	}

	return nil
}

func (a *SqlIsolationLevelSetter) SetRepeatableRead(tx common.TxController) error {
	_, err := tx.(*sql.Tx).Exec("set transaction isolation level repeatable read")
	if err != nil {
		return err
	}

	return nil
}

func (a *SqlIsolationLevelSetter) SetSerializable(tx common.TxController) error {
	_, err := tx.(*sql.Tx).Exec("set transaction isolation level serializable")
	if err != nil {
		return err
	}

	return nil
}
