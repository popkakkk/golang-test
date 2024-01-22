package liba

type LibAAllower struct {
	A bool
}

func NewLibA(a bool) LibAAllower {
	return LibAAllower{
		A: a,
	}
}

func (l LibAAllower) Allow() bool {
	return l.A
}
