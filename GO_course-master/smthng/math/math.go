package math

type Set struct {
	values  []int
	isEmpty bool
}

func New(nums ...int) Set {
	s := Set{}
	for _, v := range nums {
		s.values = append(s.values, v)
	}
	if len(s.values) > 0 {
		s.isEmpty = false
	}
	return s
}

func (s *Set) AddToSet(vals ...int) {
	for _, v := range vals {
		s.values = append(s.values, v)
	}
}

func (s *Set) IsEmpty() bool {
	return s.isEmpty
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
