package arraylist

/*
An ArrayList is a dynamically sized array that doubles capacity when it hits the limit.
It also includes a bunch of "quality of life" methods to manipulate and access it's elements.
*/
type ArrayList[T any] struct {
	length   int
	capacity int
	array    *[]T
}

func NewArrayList[T any](initialCapacity int) *ArrayList[T] {
	// Go is so strictly typed that it's impossible to initialize
	// a raw array with size as input of the function
	// arr := new([initialCapacity]T)
	// nor is it compatible with type *[]T because for arrays - size is part of the type 🤯
	arr := make([]T, initialCapacity) // so I'll use a slice for the sake of the exercise

	return &ArrayList[T]{
		length:   0,
		capacity: initialCapacity,
		array:    &arr,
	}
}

func (list *ArrayList[T]) doubleCapacity() {
	list.capacity = list.capacity * 2

	newArr := make([]T, list.capacity)
	for i := 0; i < list.length; i++ {
		newArr[i] = (*list.array)[i]
	}

	list.array = &newArr
}

func (list *ArrayList[T]) Append(item T) {
	if list.length == list.capacity {
		list.doubleCapacity()
	}

	(*list.array)[list.length] = item
	list.length++
}
