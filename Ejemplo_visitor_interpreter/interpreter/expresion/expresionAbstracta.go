package expresion

import "Ejemplo_visitor_interpreter/interpreter/tipos"

type ExpresionAbstracta interface {
	GetTipo() tipos.TipoAbstracto
	GetValor() string
}