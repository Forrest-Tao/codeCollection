package workqueue

type Interface typedInterface[t]

type typedInterface[T comparable] interface {
	Add(item T)
	Len() int
	Get() T
	Done(item T)
	ShutDown()
	ShutDownWithDrain()
	ShuttingDown() bool
}

type Queue[T comparable] interface {
	Touch(item T)
	Push(item T)
	Len() int
	Pop() T
}

func DefaultQueue[T comparable]() Queue[T] { return }

// queue is a slice which implements the Queue interface
type queue[T comparable] []T

func (q *queue[T]) Touch(item T) {}

func (q *queue[T]) Push(item T) {
	*q = append(*q, item)
}

func (q *queue[T]) Len() int { return len(*q) }

func (q *queue[T]) Pop() (item T) {
	item = (*q)[0]
	(*q)[0] = *new(T)
	*q = (*q)[1:]
	return item
}

type TypedQueueConfig[T comparable] struct {
	Name string
}

type empty struct{}

type t interface{}
type set[t comparable] map[t]empty

func (s set[t]) has(item t) bool {
	_, ok := s[item]
	return ok
}

func (s set[t]) insert(item t) {
	s[item] = empty{}
}

func (s set[t]) delete(item t) {
	delete(s, item)
}

func (s set[t]) len() int { return len(s) }
