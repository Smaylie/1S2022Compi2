package tipos

type Tipo struct {
	Nombre string
}

func NewTipo(nombre string) Tipo {
	return Tipo{Nombre: nombre}
}

func (t *Tipo) ToString() string {
	return t.Nombre
}
