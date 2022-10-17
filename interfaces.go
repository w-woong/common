package common

type Starter interface {
	Start() error
}

type Stopper interface {
	Stop() error
}
