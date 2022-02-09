package service

type Link interface {
	ID() int
	Write(data []byte) error
	//Read(data []byte) (int, error)
	Close() error
	OnClose(fn func())
	OnData(fn func(data []byte))
}
