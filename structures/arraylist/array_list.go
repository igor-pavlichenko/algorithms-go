package arraylist

import (
	"errors"
	"fmt"
)

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

func (list *ArrayList[T]) InsertAt(item T, idx int) (bool, error) {
	if list.length == list.capacity {
		list.doubleCapacity()
	}

	if idx > list.length {
		return false, errors.New("index out of bounds")
	}

	// first item insert at index 0
	if list.length == 0 && idx == 0 {
		(*list.array)[0] = item
		list.length++
		return true, nil
	}

	// otherwise we need to shift elements to make space for new item at that index
	i := list.length
	for ; i > idx; i-- {
		// keep shifting elements to the end until we reach the provided index
		(*list.array)[i] = (*list.array)[i-1]
	}

	// we reached that index, and all elements have been shifted
	// it's safe to insert the item here
	(*list.array)[i] = item
	list.length++

	return true, nil
}

func (list *ArrayList[T]) Append(item T) {
	if list.length == list.capacity {
		list.doubleCapacity()
	}

	(*list.array)[list.length] = item
	list.length++
}

func (list *ArrayList[T]) Get(index int) (T, error) {
	if list.length == 0 {
		var nilValue T
		return nilValue, errors.New("list empty")
	}
	if index < 0 || index >= list.length {
		var nilValue T
		return nilValue, errors.New("invalid index")
	}

	return (*list.array)[index], nil

}

func (list *ArrayList[T]) ToString() string {
	var str string = ""
	for i := 0; i < list.length; i++ {
		str += "["
		str += fmt.Sprint((*list.array)[i])
		str += "]"

		if i+1 < list.capacity {
			str += " - "
		}
	}

	// print empty indexes
	for j := list.length - 1; j < list.capacity; j++ {
		str += "[ ]"

		if j+1 < list.capacity {
			str += " - "
		}
	}

	return str
}
