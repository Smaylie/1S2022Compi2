package expresion

import "Ejemplo_visitor_interpreter/interpreter/tipos"

type ExpresionLiteral struct {
	Tipo  tipos.TipoAbstracto
	Valor string
}

func NewExpresionLiteral(tipo tipos.TipoAbstracto, valor string) ExpresionLiteral {
	return ExpresionLiteral{Tipo: tipo, Valor: valor}
}

func (e *ExpresionLiteral) GetTipo() tipos.TipoAbstracto {
	return e.Tipo
}

func (e *ExpresionLiteral) GetValor() string {
	return e.Valor
}
