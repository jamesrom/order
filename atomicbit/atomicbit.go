package atomicbit

import (
	"sync/atomic"
)

type Bit uint32

func New(val bool) (b *Bit) {
	b = new(Bit)
	b.Set(val)
	return b
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
