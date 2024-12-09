package quack

// queue and stack

type Quack[T any] []T

type Stack = Quack[any]
type Queue = Quack[any]

func (sq *Quack[T]) push(n T) {
	*sq = append(*sq, n)
}

func (sq *Quack[T]) pull() T {
	n := (*sq)[0]
	*sq = (*sq)[1:]
	return n
}

func (sq *Quack[T]) pop() T {
	n := (*sq)[len(*sq)-1]
	*sq = (*sq)[:len(*sq)-1]
	return n
}

func (sq *Quack[T]) first() T {
	return (*sq)[0]
}

func (sq *Quack[T]) last() T {
	return (*sq)[len(*sq)-1]
}

func (sq *Quack[T]) size() int {
	return len(*sq)
}
