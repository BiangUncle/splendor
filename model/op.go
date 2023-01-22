package model

type OP interface {
	Get(idx int, num int)
	Put(idx int, num int)
}
