package common

//go:generate mockgen -destination=./mocks/mock_tx.go -package=mocks -mock_names=TxBeginner=MockTxBeginner,TxController=MockTxController -source=./tx.go . TxBeginner,TxController

type TxBeginner interface {
	Begin() (TxController, error)
}

type RWTxBeginner interface {
	Begin() (TxController, error)
	BeginR() (TxController, error)
}

type TxController interface {
	Commit() error
	Rollback() error
}

type IsolationLevelSetter interface {
	SetReadUncommitted() error
	SetReadCommitted() error
	SetRepeatableRead() error
	SetSerializable() error
}
