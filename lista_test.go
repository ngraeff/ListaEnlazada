package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test Clasicos TDA

// Test lista vacia
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

// Test insertar al principio y final en Lista
func TestInsertarPrimeroYUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(-56)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 10, lista.VerPrimero())
	require.EqualValues(t, -56, lista.VerUltimo())
}

// Test borrar primero de Lista
func TestListaBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("hola")
	lista.InsertarPrimero("como")
	lista.InsertarPrimero("estan?")
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, "estan?", lista.BorrarPrimero())
	require.EqualValues(t, "como", lista.VerPrimero())
	lista.BorrarPrimero()
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, "hola", lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

// Test volumen Lista
func TestVolumenLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valorTope := 1000
	for i := 0; i <= valorTope; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, 1001, lista.Largo())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 1000, lista.VerUltimo())
	for i := 1000; i >= 0; i-- {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}

//------------------------------------------------------------------------
// Tests Iterador Externo

// Test iterador vacio
func TestIteradorExternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	iter.Insertar(2)

	require.EqualValues(t, 2, iter.VerActual())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, iter.Borrar())
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

// Test inserta a la ultima posicion con iterador
func TestIteradorUltimoConListaElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(4)
	require.Equal(t, 4, lista.VerUltimo())
}

// Test inserta a la primera posicion con iterador
func TestAgregarElementoPrincipioConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Insertar(4)
	require.Equal(t, 4, iter.VerActual())
}

// Test borra ultimo elemento
func TestRemoverUltimoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Siguiente()
	require.Equal(t, 3, iter.Borrar())
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

// Test borra primero con iterador y observa el cambio en lista
func TestRemoverPrimeroConIteradorLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.EqualValues(t, 2, lista.VerPrimero())

	iter := lista.Iterador()
	require.EqualValues(t, 2, iter.VerActual())
	require.EqualValues(t, 2, iter.Borrar())

	require.NotEqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerPrimero())
}

// Test borra ultimo con iterador y observa el cambio en lista
func TestRemoverUltimoConIteradorLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.EqualValues(t, 2, lista.VerPrimero())

	iter := lista.Iterador()
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())
	require.EqualValues(t, 3, iter.Borrar())

	require.NotEqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerPrimero())
}

// Test borra todos con iterador desde el principio y observa que la lista este vacia
func TestBorrarTodoPrimerosIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	iter := lista.Iterador()
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 3, iter.Borrar())
	require.Equal(t, 4, iter.Borrar())

	require.False(t, lista.Iterador().HaySiguiente())
}

// Test volumen iterador
func TestVolumenIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	valorTope := 1000
	for i := 0; i <= valorTope; i++ {
		iter.Insertar(i)
	}
	for i := 1000; i >= 0; i-- {
		require.Equal(t, i, iter.Borrar())
	}
	require.False(t, iter.HaySiguiente())
}

// Test iterador con funcion visitar
func TestIterarFuncVisitar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	i := 1
	visitar := func(dato int) bool {
		require.Equal(t, i, dato)
		i++
		return true
	}
	lista.Iterar(visitar)
}

func TestIterInsertarAlFinalListaNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(4)
	iter := lista.Iterador()
	require.EqualValues(t, 2, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 4, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

	lista.InsertarUltimo(5)
	require.EqualValues(t, 5, lista.VerUltimo())

	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 5, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

}

func TestIterPruebaFalla(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	require.EqualValues(t, 3, iter.VerActual())
	iter.Insertar(6)
	require.EqualValues(t, 6, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())

	iter2 := lista.Iterador()
	iter2.Siguiente()
	iter2.Siguiente()
	iter2.Siguiente()
	iter2.Siguiente()
	iter2.Siguiente()

}
