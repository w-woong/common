package txcom

import "gorm.io/gorm"

type GormIsolationLevelSetter struct {
	db *gorm.DB
}

func NewGormIsolationLevelSetter(db *gorm.DB) *GormIsolationLevelSetter {
	return &GormIsolationLevelSetter{
		db: db,
	}
}

func (a *GormIsolationLevelSetter) SetReadUncommitted() error {
	res := a.db.Exec("set transaction isolation level read uncommitted")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (a *GormIsolationLevelSetter) SetReadCommitted() error {
	res := a.db.Exec("set transaction isolation level read committed")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (a *GormIsolationLevelSetter) SetRepeatableRead() error {
	res := a.db.Exec("set transaction isolation level repeatable read")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (a *GormIsolationLevelSetter) SetSerializable() error {
	res := a.db.Exec("set transaction isolation level serializable")
	if res.Error != nil {
		return res.Error
	}

	return nil
}
