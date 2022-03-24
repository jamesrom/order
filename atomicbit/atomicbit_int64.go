//go:build atomicbit_int64
// +build atomicbit_int64

package atomicbit

import (
	"sync/atomic"
)

type Bit int64

func New(val bool) (b *Bit) {
	i := int64(0)
	if val {
		i = 1
	}
	return (*Bit)(&i)
}

func (b *Bit) Flip() {
	atomic.AddInt64((*int64)(b), 1)
}

func (b *Bit) Get() bool {
	return atomic.LoadInt64((*int64)(b))%2 == 1
}

func (b *Bit) Set(val bool) {
	if val {
		atomic.StoreInt64((*int64)(b), 1)
	} else {
		atomic.StoreInt64((*int64)(b), 0)
	}
}
