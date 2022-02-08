package service

type Link interface {
	Write(data []byte) error
	Read(data []byte) error
	Close() error
}