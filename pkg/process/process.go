package process

type ProcessInterface interface {
	Process(opts any) (interface{}, error)
	Stringify(opts any) string
}
