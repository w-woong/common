package txcom

import (
	"github.com/w-woong/common"
)

type GormIsolationLevelSetter struct {
}

func NewGormIsolationLevelSetter() *GormIsolationLevelSetter {
	return &GormIsolationLevelSetter{}
}

func (a *GormIsolationLevelSetter) SetReadUncommitted(tx common.TxController) error {
	res := tx.(*GormTxController).Tx.Exec("set transaction isolation level read uncommitted")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (a *GormIsolationLevelSetter) SetReadCommitted(tx common.TxController) error {
	res := tx.(*GormTxController).Tx.Exec("set transaction isolation level read committed")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (a *GormIsolationLevelSetter) SetRepeatableRead(tx common.TxController) error {
	res := tx.(*GormTxController).Tx.Exec("set transaction isolation level repeatable read")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (a *GormIsolationLevelSetter) SetSerializable(tx common.TxController) error {
	res := tx.(*GormTxController).Tx.Exec("set transaction isolation level serializable")
	if res.Error != nil {
		return res.Error
	}

	return nil
}
