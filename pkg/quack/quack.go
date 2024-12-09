package quack

// queue and stack

type Quack[T any] []any

func NewStack[T any]() *Quack[T] {
	return &Quack[T]{}
}

func NewQueue[T any]() *Quack[T] {
	return &Quack[T]{}
}

func (sq *Quack[T]) push(n any) {
	*sq = append(*sq, n)
}

func (sq *Quack[T]) pull() any {
	n := (*sq)[0]
	*sq = (*sq)[1:]
	return n
}

func (sq *Quack[T]) pop() any {
	n := (*sq)[len(*sq)-1]
	*sq = (*sq)[:len(*sq)-1]
	return n
}

func (sq *Quack[T]) first() any {
	return (*sq)[0]
}

func (sq *Quack[T]) last() any {
	return (*sq)[len(*sq)-1]
}

func (sq *Quack[T]) size() any {
	return len(*sq)
}
