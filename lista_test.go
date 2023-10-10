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
		require.EqualValues(t, i, lista.VerUltimo())
		require.EqualValues(t, 0, lista.VerPrimero())
	}
	require.EqualValues(t, 1001, lista.Largo())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 1000, lista.VerUltimo())
	for i := 0; i <= 1000; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

//------------------------------------------------------------------------
// Tests Iterador Externo

// Test iterador vacio
func TestIteradorExternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })

	iterador.Insertar(2)

	require.EqualValues(t, 2, iterador.VerActual())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, iterador.Borrar())
	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}

func TestIterExternoBasico(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(30)
	lista.InsertarPrimero(40)
	iterador := lista.Iterador()
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, 40, iterador.VerActual())
	iterador.Siguiente()
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, 30, iterador.VerActual())
	iterador.Siguiente()
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, 20, iterador.VerActual())
	iterador.Siguiente()
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, 10, iterador.VerActual())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())
}

// Test inserta a la ultima posicion con iterador
func TestIteradorUltimoConListaElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Insertar(4)
	require.Equal(t, 4, lista.VerUltimo())
}

// Test inserta a la primera posicion con iterador
func TestAgregarElementoPrincipioConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	iterador.Insertar(4)
	require.Equal(t, 4, iterador.VerActual())
}

// Test borra ultimo elemento
func TestRemoverUltimoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	iterador.Siguiente()
	require.Equal(t, 3, iterador.Borrar())
	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}
func TestRemoverPrimeroConIteradorLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.EqualValues(t, 2, lista.VerPrimero())

	iterador := lista.Iterador()
	require.EqualValues(t, 2, iterador.VerActual())
	require.EqualValues(t, 2, iterador.Borrar())

	require.NotEqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerPrimero())
}

// Test borra ultimo con iterador y observa el cambio en lista
func TestRemoverUltimoConIteradorLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.EqualValues(t, 2, lista.VerPrimero())

	iterador := lista.Iterador()
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	require.EqualValues(t, 3, iterador.Borrar())

	require.NotEqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerPrimero())
}

// Test borra todos con iterador desde el principio y observa que la lista este vacia
func TestBorrarTodoPrimerosIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	iterador := lista.Iterador()
	require.Equal(t, 2, iterador.Borrar())
	require.Equal(t, 3, iterador.Borrar())
	require.Equal(t, 4, iterador.Borrar())

	require.False(t, lista.Iterador().HaySiguiente())
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

// Test iterador con funcion buscar
func TestIterarFuncBuscar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	target := 1
	buscar := func(dato int) bool {
		if target == dato {
			require.Equal(t, target, dato)
			return false
		}
		return true
	}

	lista.Iterar(buscar)
}

func TestIterInsertarAlFinalListaNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(4)
	iterador := lista.Iterador()
	require.EqualValues(t, 2, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 4, iterador.VerActual())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())

	lista.InsertarUltimo(5)
	require.EqualValues(t, 5, lista.VerUltimo())

	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, 5, iterador.VerActual())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())

}

func TestIterPruebaFalla(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()

	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Insertar(6)
	require.EqualValues(t, 6, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())

	iterador2 := lista.Iterador()
	iterador2.Siguiente()
	iterador2.Siguiente()
	iterador2.Siguiente()
	iterador2.Siguiente()
	iterador2.Siguiente()

}

// Test volumen iterador
func TestVolumenIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	valorTope := 1000
	for i := 0; i <= valorTope; i++ {
		iterador.Insertar(i)
	}
	for i := 1000; i >= 0; i-- {
		require.Equal(t, i, iterador.Borrar())
	}
	require.False(t, iterador.HaySiguiente())
}

// Test volumen iterador Externo
func TestVolumenIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	valorTope := 1000

	for i := 0; i <= valorTope; i++ {
		iterador.Insertar(i)
		require.EqualValues(t, i, lista.VerPrimero())
	}
	iterador2 := lista.Iterador()
	for i:=valorTope;i>=0;i--{
		require.EqualValues(t,i,iterador2.VerActual())
		iterador2.Siguiente()
	}
	require.EqualValues(t, 1001, lista.Largo())
	require.EqualValues(t, 1000, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())

	for i := 1000; i >= 0; i-- {
		require.EqualValues(t, i, iterador.Borrar())
	}
	require.True(t, lista.EstaVacia())
}

// Test iterador interno con gran volumen
func TestIterarInternoVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valorTope := 1000
	for i := 0; i <= valorTope; i++ {
		lista.InsertarUltimo(i)
	}

	target := 10
	buscar := func(dato int) bool {
		if target == dato {
			return false
		}
		return true
	}

	lista.Iterar(buscar)
}
