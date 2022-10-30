package txcom

import (
	"sync"

	"github.com/w-woong/common"
)

type LockTxBeginner struct {
	l *sync.RWMutex
}

func NewLockTxBeginner() *LockTxBeginner {
	return &LockTxBeginner{
		l: &sync.RWMutex{},
	}
}

func (a *LockTxBeginner) Begin() (common.TxController, error) {
	a.l.Lock()
	return NewLockTxController(a.l), nil
}

func (a *LockTxBeginner) BeginR() (common.TxController, error) {
	a.l.RLock()
	return NewRLockTxController(a.l), nil
}

type LockTxController struct {
	unlocked bool
	l        *sync.RWMutex
}

func NewLockTxController(l *sync.RWMutex) *LockTxController {
	return &LockTxController{
		unlocked: false,
		l:        l,
	}
}

func (a *LockTxController) Commit() error {
	if a.unlocked {
		return nil
	}
	a.unlocked = true
	a.l.Unlock()
	return nil
}

func (a *LockTxController) Rollback() error {
	if a.unlocked {
		return nil
	}
	a.unlocked = true
	a.l.Unlock()
	return nil
}

type RLockTxController struct {
	unlocked bool
	l        *sync.RWMutex
}

func NewRLockTxController(l *sync.RWMutex) *RLockTxController {
	return &RLockTxController{
		unlocked: false,
		l:        l,
	}
}

func (a *RLockTxController) Commit() error {
	if a.unlocked {
		return nil
	}
	a.unlocked = true
	a.l.RUnlock()
	return nil
}

func (a *RLockTxController) Rollback() error {
	if a.unlocked {
		return nil
	}
	a.unlocked = true
	a.l.RUnlock()
	return nil
}
