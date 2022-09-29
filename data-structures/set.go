package datastructures

type Set[T comparable] struct {
	data map[T]struct{}
	size int
}

func CreateSet[T comparable]() Set[T] {
	s := Set[T]{
		data: make(map[T]struct{}),
		size: 0,
	}

	return s
}

func (s *Set[T]) Add(val T) {
	s.data[val] = struct{}{}
}

func (s *Set[T]) Remove(val T) {
	delete(s.data, val)
}

func (s *Set[T]) Exists(val T) bool {
	_, ok := s.data[val]
	return ok
}

func (s *Set[T]) GetValues() []T {
	keys := make([]T, 0, len(s.data))

	for k := range s.data {
		keys = append(keys, k)
	}

	return keys
}

func (s *Set[T]) Size() int {
	s.calculateSize()
	return s.size
}

func (s *Set[T]) calculateSize() {
	s.size = len(s.data)
}
