package interfaces

type DataSource interface {
	Name() string
	Publish() error
	Subscribe(callback func(string, string)) error
	Get(method string, param interface{}) (string, string)
	Close() error
}
