package queue

type (
	// Queue - Очередь представляется в качестве линейного списка,
	// в котором добавление/удаление элементов идет строго с соответствующих его концов.
	Queue[E any] struct {
		start  *node[E]
		end    *node[E]
		length int
	}
	node[E any] struct {
		value *E
		next  *node[E]
	}
)

// New - создание новой очереди
func New[E any]() *Queue[E] {
	return &Queue[E]{}
}

// Dequeue - Удаление элемента из передней части очереди и возврат его значения.
func (que *Queue[E]) Dequeue() *E {
	if que.length == 0 {
		return nil
	}
	n := que.start
	if que.length == 1 {
		que.start = nil
		que.end = nil
	} else {
		que.start = que.start.next
	}
	que.length--
	return n.value
}

// Enqueue - Добавить новый элемент в конец очереди.
func (que *Queue[E]) Enqueue(value E) {
	n := &node[E]{&value, nil}
	if que.length == 0 {
		que.start = n
		que.end = n
	} else {
		que.end.next = n
		que.end = n
	}
	que.length++
}

// Len - Возвращает количество элементов внутри очереди.
func (que *Queue[E]) Len() int {
	return que.length
}

// Last - Вернуть значение элемента в начале очереди, не удаляя его
func (que *Queue[E]) Last() *E {
	if que.length == 0 {
		return nil
	}
	return que.start.value
}
