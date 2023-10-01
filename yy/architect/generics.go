package main

func Contains [T comparable] (needle T, haystack []T) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

type Set [T comparable] struct {
	values	map[T]struct{}
}

func NewSet [T comparable] () *Set[T] {
	return &Set[T]{
		values: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(values ...T) {
	for _, value := range values {
		s.values[value] = struct{}{}
	} 
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}