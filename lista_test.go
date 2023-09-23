package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "tdas/lista"
	"testing"
)

//Test Clasicos TDA
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t,"La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsertarPrimeroYUltimo(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(-56)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t,10, lista.VerPrimero())
	require.EqualValues(t,-56 , lista.VerUltimo())
}

func TestBorrarPrimero(t *testing.T){
	lista := TDALista.CrearListaEnlazada[string]()	
	lista.InsertarPrimero("hola")
	lista.InsertarPrimero("como")
	lista.InsertarPrimero("estan?")
	require.EqualValues(t,3,lista.Largo())
	require.EqualValues(t,"estan?",lista.BorrarPrimero())
	require.EqualValues(t,"como", lista.VerPrimero())
	lista.BorrarPrimero()
	require.EqualValues(t,1,lista.Largo())
	require.EqualValues(t,"hola",lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestVolumenLista(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	valorTope :=  1000
	for i := 0; i <= valorTope; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, 1001, lista.Largo())
	require.EqualValues(t,0,lista.VerPrimero())
	require.EqualValues(t,1000,lista.VerUltimo())
	for i := 1000; i >=0; i-- {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}

//Test Iterador Externo