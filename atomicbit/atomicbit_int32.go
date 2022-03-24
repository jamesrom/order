//go:build atomicbit_int32
// +build atomicbit_int32

package atomicbit

import (
	"sync/atomic"
)

type Bit int32

func New(val bool) (b *Bit) {
	i := int32(0)
	if val {
		i = 1
	}
	return (*Bit)(&i)
}

func (b *Bit) Flip() {
	atomic.AddInt32((*int32)(b), 1)
}

func (b *Bit) Get() bool {
	return atomic.LoadInt32((*int32)(b))%2 == 1
}

func (b *Bit) Set(val bool) {
	if val {
		atomic.StoreInt32((*int32)(b), 1)
	} else {
		atomic.StoreInt32((*int32)(b), 0)
	}
}
