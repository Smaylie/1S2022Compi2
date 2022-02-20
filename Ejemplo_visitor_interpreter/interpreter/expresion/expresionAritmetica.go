package expresion

import (
	"Ejemplo_visitor_interpreter/interpreter/tipos"
	"fmt"
	"strconv"
)

const (
	MAS int = 0
	MEN     = 1
	POR     = 2
	DIV     = 3
)

type ExpresionAritmetica struct {
	Izq ExpresionAbstracta
	Der ExpresionAbstracta
	Ope int
}

func NewExpresionAritmetica(izq ExpresionAbstracta, ope int, der ExpresionAbstracta) ExpresionAritmetica {
	return ExpresionAritmetica{Izq: izq, Ope: ope, Der: der}
}

func (e *ExpresionAritmetica) GetTipo() tipos.TipoAbstracto {
	t1 := e.Izq.GetTipo()
	t2 := e.Der.GetTipo()
	return max(t1, t2)
}

func (e *ExpresionAritmetica) GetValor() string {
	tr := e.GetTipo().ToString()
	cad1 := e.Izq.GetValor()
	cad2 := e.Der.GetValor()

	if tr == tipos.NUMBER {
		v1, _ := strconv.ParseFloat(cad1, 64)
		v2, _ := strconv.ParseFloat(cad2, 64)
		switch e.Ope {
		case MAS:
			return fmt.Sprintf("%.2f", v1+v2)
		case MEN:
			return fmt.Sprintf("%.2f", v1-v2)
		case POR:
			return fmt.Sprintf("%.2f", v1*v2)
		default:
			return fmt.Sprintf("%.2f", v1/v2)
		}
	} else if tr == tipos.STRING {
		if e.Ope == MAS {
			return cad1 + cad2
		} else {
			// error
			fmt.Println("error al operar cadenas")
			return ""
		}
	} else {
		// error
		return ""
	}
}

func max(t1 tipos.TipoAbstracto, t2 tipos.TipoAbstracto) tipos.TipoAbstracto {
	if t1.ToString() == tipos.STRING || t2.ToString() == tipos.STRING {
		nt := tipos.NewTipo(tipos.STRING)
		return &nt
	} else if t1.ToString() == tipos.NUMBER || t2.ToString() == tipos.NUMBER {
		nt := tipos.NewTipo(tipos.NUMBER)
		return &nt
	} else {
		nt := tipos.NewTipo(tipos.ERROR)
		return &nt
	}
}
