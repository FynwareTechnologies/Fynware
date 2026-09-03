package devices

type Device interface {
	Name() string
	Reset()
	Start() error
	Stop() error
}
