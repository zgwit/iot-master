package service

type Service interface {
	Open() error
	Close() error
	HasAcceptor() bool
	GetLink(id int)(Link, error)
}