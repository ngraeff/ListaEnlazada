package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterador[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

// CrearListaEnlazada crea una nueva lista enlazada vacia.
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

// crearNodo Crea un nuevo nodo con un dato.
func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato, nil}
}

// Largo devuelve la cantidad de elementos de la lista.
func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

// EstaVacia devuelve verdadero si la lista no tiene elementos encolados, false en caso contrario.
func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.Largo() == 0
}

// InsertarPrimero agrega un nuevo elemento a le lista en el primer lugar.

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo(dato)
	if l.EstaVacia() {
		l.ultimo = nodo
	} else {
		nodo.prox = l.primero
	}
	l.primero = nodo
	l.largo++
}

// InsertarUltimo agrega un nuevo elemento a le lista en el último lugar.
func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodo(dato)

	if l.EstaVacia() {
		l.primero = nodo
	} else {
		l.ultimo.prox = nodo
	}
	l.ultimo = nodo
	l.largo++
}

// VerPrimero devuelve el valor del primer elemento de la lista. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

// VerUltimo devuelve el valor del ultimo elemento de la lista.Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

// BorrarPrimero borra el primer elemento de la lista y devuelve ese valor.  Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	prime := l.primero
	l.primero = prime.prox
	l.largo--

	return prime.dato
}

// Iterar pasa por cada uno de los elementos de la lista en orden hasta acabarla
func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		continuar := visitar(actual.dato)
		if !continuar {
			break
		}
		actual = actual.prox
	}
}

// Iterador crea un iterador de la lista externo.
func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterador[T]{l, l.primero, nil}
}

// VerActual devuelve el elemento en donde este posicionado el iterador.
func (i *iterador[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.actual.dato
}

// Siguiente modifica la posicion del iterador al siguiente elemento.
func (i *iterador[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	i.anterior = i.actual
	i.actual = i.actual.prox

}

// HaySiguiente devuelve verdadero si todavía hay algun elemento de la lista por ver, en caso contrario devuelve falso
func (i *iterador[T]) HaySiguiente() bool {
	if i.actual == nil && i.anterior != i.lista.ultimo {
		i.actual = i.lista.ultimo
		return true
	}
	return i.actual != nil
}

// Insertar inserta un elemento a la lista en donde este posicionado el iterador.
func (i *iterador[T]) Insertar(dato T) {
	nodo := crearNodo(dato)
	nodo.prox = i.actual
	if i.anterior == nil {
		i.lista.primero = nodo
	}
	if !i.HaySiguiente() {
		i.lista.ultimo = nodo
	}
	if i.anterior != nil {
		i.anterior.prox = nodo
	}

	i.actual = nodo
	i.lista.largo++
}

// Borrar borra el elemento de la lista en donde este posicionado el iterador, y ademas devuelve el valor de ese elemento.
func (i *iterador[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := i.VerActual()

	if i.anterior == nil { //Primer Elemento
		i.lista.primero = i.actual.prox
	}
	if i.actual.prox == nil { //Ultimo Elemento
		i.lista.ultimo = i.anterior
	}

	if i.anterior != nil { //Cualquier Elemento excepto el primero
		i.anterior.prox = i.actual.prox
	}
	i.actual = i.actual.prox

	i.lista.largo--
	return dato
}
