package lista_test

import (
	"fmt"
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

//TEST PARA LISTA EN GENERAL

func TestListaCreadaEstaVacia(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
}

func TestListsCreadaPanics(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

}

func TestInsertarPrimeroLista(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	lista.InsertarPrimero(4)
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia(), "La lista esta vacia si el largo es 0")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")



}

func TestInsertarUltimoLista(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarUltimo(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())
	lista.InsertarUltimo(3)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	lista.InsertarUltimo(4)
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia(), "La lista esta vacia si el largo es 0")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")
}

func TestVaciarLista(t *testing.T){
	lista := TDALista.CrearListaEnlazada(int)

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	
	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	lista.InsertarUltimo(3)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())

	lista.InsertarUltimo(5)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())

	lista.InsertarUltimo(6)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 6, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 5, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 6, lista.BorrarPrimer())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.BorrarPrimer())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 1, lista.BorrarPrimer())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 3, lista.BorrarPrimer())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 5, lista.BorrarPrimer())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia si el largo es 0")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

}

func TestListarEnteros(t *testing.T){
	lista := TDALista.CrearListaEnlazada(int)

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo(0)
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, 0, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía si el largo es 0")
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")
	
}

func TestListarStrings(t *testing.T){
	lista := TDALista.CrearListaEnlazada(string)

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero("uno")
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, "uno", lista.VerPrimero())
	require.Equal(t, "uno", lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo("dos")
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, "uno", lista.VerPrimero())
	require.Equal(t, "dos", lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, "uno", lista.BorrarPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, "dos", lista.VerPrimero())
	require.Equal(t, "dos", lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, "uno", lista.BorrarPrimero())
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía si el largo es 0")
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")
}

func TestListarMuchosElementos(t *testing.T){
	lista := TDALista.CrearListaEnlazada(int)

	const n = 10000
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())
	for(i:= 0; i < n; i++){
		lista.InsertarPrimero(i)
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 0, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}

	require.Equal(t, n, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")

	for(i:= lista.Largo()-1; i > 0; i--){
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, i, lista.BorrarPrimer())
		require.Equal(t, i-1, lista.VerPrimero())
		require.Equal(t, 0, lista.VerUltimo())
		require.Equal(t, i-1, lista.Largo())
	}

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())
	for(i:= 0; i < n; i++){
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, 0, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}

	require.Equal(t, n, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")

	for(i:= 0; i < n; i++){
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, i, lista.BorrarPrimer())
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, n-1, lista.VerUltimo())
		require.Equal(t, i-1, lista.Largo())
	}

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

}

func TestInsertarYBorrarIntercalado(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.Largo())

	lista.InsertarUltimo(5)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())

	lista.InsertarPrimero(7)
	require.Equal(t, 7, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 7, lista.BorrarPrimero())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())

	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 4, lista.Largo())

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, 2, lista.BorrarPrimero())
	
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía si su largo es 0")
	require.Equal(t, 0, lista.Largo())
	
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")


}

//TEST PARA ITERADORES:

//EXTERNO

//func TestIteradorExtCreado (t *testing.T)

//func TestIterarExtIterarListaVacia(t *testing.T)

//func TestIteradorExtInsertarYBorrarUno(t *testing.T)

//func TestIteradorExtInsertarPrimero(t *testing.T)

//func TestIteradorExtInsertarUltimo(t *testing.T)

//func TestIteradorExtInsertarMedio(t *testing.T)

//func TestIteradorExtIteracionCompleta(t *testing.T)

//func TestIteradorExtBorrarPrimerElemento(t *testing.T)

//func TestIteradorExtBorrarUltimoElemento(t *testing.T)

//func TestIteradorExtBorrarElementoMedio(t *testing.T)

//INTERNO

//func TestIteradorIntIterarTodos(t *testing.T)

//func TestIterarIntIterarListaVacia(t *testing.T)

//func TestIterarIntIterarConCorte(t *testing.T)
