package main

import (
	"context"
	"time"
)

type callBackChan chan struct{}

func checkEvery(ctx context.Context, d time.Duration, cb callBackChan) {
	for {
		select {
		case <-ctx.Done():
			// ctx is canceled
			return
		case <-time.After(d):
			// wait for the duration
			if cb != nil {
				cb <- struct{}{}
			}
		}
	}
}

func main() {

}
