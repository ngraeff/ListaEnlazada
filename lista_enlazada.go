package  lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo int
}

type iterador[T any] struct{
	actual *nodoLista[T]
	anterior *nodoLista[T]
	largo int
}

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

// CrearListaEnlazada crea una nueva lista enlazada vacia.
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil,0}
}

//crearNodo Crea un nuevo nodo con un dato.
func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato, nil}
}

// Largo devuelve la cantidad de elementos de la lista.
func (l *listaEnlazada)Largo() int{
	return l.largo
}

// EstaVacia devuelve verdadero si la lista no tiene elementos encolados, false en caso contrario.
func (l *listaEnlazada)EstaVacia() bool{
	return l.Largo() == 0
}

// InsertarPrimero agrega un nuevo elemento a le lista en el primer lugar.

func (l *listaEnlazada)InsertarPrimero(dato T){
	nodo := crearNodo(dato)
	if l.estaVacia(){
		l.ultimo = nodo
	} else{
		nodo.prox= l.primero
	}
	l.primero = nodo
	l.largo++
}

// InsertarUltimo agrega un nuevo elemento a le lista en el último lugar.
func (l *listaEnlazada)InsertarUltimo(dato T){
	nodo := crearNodo(dato)

	if l.estaVacia(){
		l.primero = nodo
	}else{
		l.ultimo.prox = nodo
	}
	l.ultimo = nodo
	l.largo ++
}

// VerPrimero devuelve el valor del primer elemento de la lista. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (l *listaEnlazada) VerPrimero() T{
	if l.estaVacia(){
		panic("La lista esta Vacía")
	}
	return l.primero.dato
}

// VerUltimo devuelve el valor del ultimo elemento de la lista.Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (l *listaEnlazada) VerUltimo() T{
	if l.estaVacia(){
		panic("La lista esta Vacía")
	}
	return l.ultimo.dato
}

// BorrarPrimero borra el primer elemento de la lista y devuelve ese valor.  Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (l *listaEnlazada) BorrarPrimero() T{
	if l.estaVacia(){
		panic("La lista esta vacia")
	}
	primero := l.VerPrimero()
	l.primero = primero.prox
	l.largo --

	return primero.dato
}

func (l *listaEnlazada) Iterador() IteradorLista[T]{
	if l.EstaVacia(){
		primero := nil
	}else{
		primero:= l.VerPrimero()
	}
	return &iterador{primero,nil,l.Largo()}
}