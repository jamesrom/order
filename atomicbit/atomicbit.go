//go:build !atomicbit_int32 && !atomicbit_uint64 && !atomicbit_int64 && !atomicbit_uintptr
// +build !atomicbit_int32,!atomicbit_uint64,!atomicbit_int64,!atomicbit_uintptr

package atomicbit

import (
	"sync/atomic"
)

type Bit uint32

func New(val bool) (b *Bit) {
	i := uint32(0)
	if val {
		i = 1
	}
	return (*Bit)(&i)
}

func (b *Bit) Flip() {
	atomic.AddUint32((*uint32)(b), 1)
}

func (b *Bit) Get() bool {
	return atomic.LoadUint32((*uint32)(b))%2 == 1
}

func (b *Bit) Set(val bool) {
	if val {
		atomic.StoreUint32((*uint32)(b), 1)
	} else {
		atomic.StoreUint32((*uint32)(b), 0)
	}
}
