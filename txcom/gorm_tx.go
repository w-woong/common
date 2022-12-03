package txcom

import (
	"errors"

	"github.com/w-woong/common"
	"gorm.io/gorm"
)

// GormTxBeginner is a transaction beginner for gorm.DB
type GormTxBeginner struct {
	db *gorm.DB
}

// NewGormTxBeginner returns a new GormTxBeginner
func NewGormTxBeginner(db *gorm.DB) *GormTxBeginner {
	return &GormTxBeginner{
		db: db,
	}
}

// Begin starts transaction returning common.TxController that commits or rollbacks
func (a *GormTxBeginner) Begin() (common.TxController, error) {
	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return NewGormTxController(tx), nil
}

// Begin starts transaction returning common.TxController that commits or rollbacks
func (a *GormTxBeginner) BeginR() (common.TxController, error) {
	return a.Begin()
}

// GormTxController is a transcation controller for gorm.DB that commits or rollbacks
type GormTxController struct {
	Tx *gorm.DB
}

// NewGormTxController returns a new GormTxController
func NewGormTxController(tx *gorm.DB) *GormTxController {
	return &GormTxController{
		Tx: tx,
	}
}

// Commit commits a transaction
func (a *GormTxController) Commit() error {
	return a.Tx.Commit().Error
}

// Rollback rollbacks and cancels a transaction
func (a *GormTxController) Rollback() error {
	return a.Tx.Rollback().Error
}

// ConvertErr converts gorm package's errors to internal ones
func ConvertErr(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.ErrRecordNotFound
	}

	return err
}
