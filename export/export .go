package export

type ExportInterface interface {
	Export() error
}

type exportMethod struct {
	method ExportInterface
}

func NewExport(method ExportInterface) *exportMethod {
	return &exportMethod{
		method: method,
	}
}

func (e exportMethod) Export() error {
	return e.method.Export()
}
