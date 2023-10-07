package stack

type (
	// Stack - список элементов
	// Зачастую стек реализуется в виде однонаправленного списка
	// (каждый элемент в списке содержит помимо хранимой информации в стеке указатель
	// на следующий элемент стека).
	Stack[E any] struct {
		top    *node[E]
		length int
	}
	node[E any] struct {
		value *E
		prev  *node[E]
	}
)

// New - создание нового стэка
func New[E any]() *Stack[E] {
	return &Stack[E]{nil, 0}
}

// Len - возврат количества элементов в стеке
func (st *Stack[E]) Len() int {
	return st.length
}

// Last - возврат верхнего элемента
func (st *Stack[E]) Last() *E {
	if st.length == 0 {
		return nil
	}
	return st.top.value
}

// Get - возврат элемента и удаление его
func (st *Stack[E]) Get() *E {
	if st.length == 0 {
		return nil
	}

	n := st.top
	st.top = n.prev
	st.length--
	return n.value
}

// Set - значение в верхней части стека
func (st *Stack[E]) Set(value E) {
	n := &node[E]{&value, st.top}
	st.top = n
	st.length++
}
