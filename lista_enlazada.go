package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type iteradorExterno[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]

	//Para acordarme (despues borrar comentario) : le pse listaItera para no confundir con lista al poner lista.alg
	listaIterar *listaEnlazada[T]
}

// func CrearListaEnlazada[T any]() Lista[T]

//func (lista *listaEnlazada[T]) EstaVacia() bool

//func (lista *listaEnlazada[T]) VerPrimero() T

//func (lista *listaEnlazada[T]) VerUltimo() T

//func (lista *listaEnlazada[T]) Largo() int

//func (lista *listaEnlazada[T]) BorrarPrimero() T

//func crearNodo[T any](datoNuevo T) *nodoLista[T]

//func (lista *listaEnlazada[T]) InsertarPrimero(dato T)

//func (lista *listaEnlazada[T]) InsertarUltimo(dato T)

// ITERADOR INTERNO

// ITERADOR EXTERNO

//func Crear iterador (en realidad es la de abajo)
//func (lista *listaEnlazada[T]) Iterador IteradorLista[T]

// PRIMITIVAS DE ITERADOR:

//func (iterador *IteradorLista[T]) VerActual() T

//func (iterador *IteradorLista[T]) HaySiguiente() bool

//func (iterador *IteradorLista[T]) Siguiente()

//func (iterador *IteradorLista[T]) Insertar()

//func (iterador *IteradorLista[T]) Borrar() T
