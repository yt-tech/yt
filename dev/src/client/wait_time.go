package client

import (
	"time"
)

type baseToken struct {
	ack      int32
	complete chan struct{}
	ready    bool
	err      error
}

// Wait will wait indefinitely for the Token to complete, ie the Publish
// to be sent and confirmed receipt from the broker
func (b *baseToken) Wait() {
	if !b.ready {
		<-b.complete
		b.ready = true
	}
}

// WaitTimeout takes a time in ms
func (b *baseToken) WaitTimeout(d time.Duration) {
	if !b.ready {
		select {
		case <-b.complete:
			b.ready = true
		case <-time.After(d):
			b.ack = 9
		}
	}
}

func (b *baseToken) flowComplete() {
	close(b.complete)
}
