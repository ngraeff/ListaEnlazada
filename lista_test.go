package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test Clasicos TDA
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsertarPrimeroYUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(-56)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 10, lista.VerPrimero())
	require.EqualValues(t, -56, lista.VerUltimo())
}

func TestBorrarPrimero(t *testing.T) {
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

// Test Iterador Externo
func TestIteradorExternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(2)

	require.EqualValues(t, 2, iter.VerActual())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, iter.Borrar())
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

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

// Test agrega un elemento al principio y se fija si se agrego correctamente
func TestAgregarElementoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Insertar(4)
	require.Equal(t, 4, iter.VerActual())
}

// Test borra un elemento del medio
func TestEliminarElementoDelMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Insertar(2)
	iter.Insertar(3)
	iter.Insertar(4)
	iter.Insertar(5)
	require.Equal(t, 5, iter.VerActual())
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 3, iter.VerActual())
	require.Equal(t, 3, iter.Borrar())
	require.Equal(t, 2, iter.VerActual())
}

// Test Iterador con funcion visitar
func TestIterarFunc(t *testing.T) {
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
