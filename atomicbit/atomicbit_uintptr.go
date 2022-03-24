//go:build atomicbit_uintptr
// +build atomicbit_uintptr

package atomicbit

import (
	"sync/atomic"
)

type Bit uintptr

func New(val bool) (b *Bit) {
	i := uintptr(0)
	if val {
		i = 1
	}
	return (*Bit)(&i)
}

func (b *Bit) Flip() {
	atomic.AddUintptr((*uintptr)(b), 1)
}

func (b *Bit) Get() bool {
	return atomic.LoadUintptr((*uintptr)(b))%2 == 1
}

func (b *Bit) Set(val bool) {
	if val {
		atomic.StoreUintptr((*uintptr)(b), 1)
	} else {
		atomic.StoreUintptr((*uintptr)(b), 0)
	}
}
