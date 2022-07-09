package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	buffer chan string
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext { //i make new context, and init a buffer in this func()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	newBuffer := make(chan string, bufferSize)
	newBufferCtx := &BufferedContext{Context: ctx, buffer: newBuffer, CancelFunc: cancel}
	return newBufferCtx
}

func (bc *BufferedContext) Done() <-chan struct{} {

	if len(bc.buffer) == cap(bc.buffer) {
		fmt.Println("Channel is full") // to check if the channel is full
		bc.CancelFunc()
	}
	return bc.Context.Done()

}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fn(bc, bc.buffer)
}

func main() {
	ctx := NewBufferedContext(3*time.Second, 5)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():

				if len(buffer) != cap(buffer) {
					fmt.Println("Channel is timedout") //added this to check if the channel is timeout
				}
				return

			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200) //if we change the sleep time, the buffer will be filled slower
				fmt.Println("bar")
			}
		}
	})
}
