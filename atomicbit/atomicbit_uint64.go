//go:build atomicbit_uint64
// +build atomicbit_uint64

package atomicbit

import (
	"sync/atomic"
)

type Bit uint64

func New(val bool) (b *Bit) {
	i := uint64(0)
	if val {
		i = 1
	}
	return (*Bit)(&i)
}

func (b *Bit) Flip() {
	atomic.AddUint64((*uint64)(b), 1)
}

func (b *Bit) Get() bool {
	return atomic.LoadUint64((*uint64)(b))%2 == 1
}

func (b *Bit) Set(val bool) {
	if val {
		atomic.StoreUint64((*uint64)(b), 1)
	} else {
		atomic.StoreUint64((*uint64)(b), 0)
	}
}
