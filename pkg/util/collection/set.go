package collection

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// Add adds the given value to the set.
func (s Set[T]) Add(value T) Set[T] {
	s[value] = struct{}{}

	return s
}

// AddAll adds all the given value to the set.
func (s Set[T]) AddAll(other Set[T]) Set[T] {
	for v := range other {
		s.Add(v)
	}

	return s
}

// Contains returns true if the given value is in the set, or false if it is not.
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) DoesNotContain(value T) bool {
	return !s.Contains(value)
}

// Remove removes the given value from the set.
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Values returns a slice of all the values in the set.
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for value := range s {
		values = append(values, value)
	}
	return values
}

func (s Set[T]) Length() int {
	return len(s)
}
