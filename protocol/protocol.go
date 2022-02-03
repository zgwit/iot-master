package protocol

type Protocol interface {
	Write(slave, code, address int, data []byte) error
	Read(slave, code, address, length int) ([]byte, error)

	// ImmediateRead 立即读，高优先级
	ImmediateRead(slave, code, address, length int) ([]byte, error)
}
