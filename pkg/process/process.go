package process

type Process interface {
	Process(data []byte) ([]byte, error)
}
