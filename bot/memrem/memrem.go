package memrem

type MR interface {
	Mem(any)
	Rem(any, int)
}

type object struct{}
