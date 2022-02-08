package service

type Service interface {
	Open() error
	Close() error
	GetLink(id int)(Link, error)
}