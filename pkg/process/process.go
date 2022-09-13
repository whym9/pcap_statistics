package process

type Process interface {
	Process(data []byte) (string, error)
}
