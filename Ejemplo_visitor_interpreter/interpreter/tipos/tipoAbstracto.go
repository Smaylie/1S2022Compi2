package tipos

const (
	NUMBER string = "number"
	STRING        = "string"
	ERROR         = "error"
)

type TipoAbstracto interface {
	ToString() string
}
