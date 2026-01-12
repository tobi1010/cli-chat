package spinner

import (
	"fmt"
	"sync"
	"time"
)

var chars = []rune{'|', '/', '-', '\\'}

const (
	hideCursor = "\x1b[?25l"
	showCursor = "\x1b[?25h"
)

func Spin() (stop func()) {
	done := make(chan struct{})
	var once sync.Once
	fmt.Print(hideCursor)
	go func() {
		i := 0
		t := time.NewTicker(100 * time.Millisecond)
		defer t.Stop()

		for {
			select {
			case <-done:
				fmt.Print("\r\n" + showCursor)
				return
			case <-t.C:
				fmt.Printf("\r%c", chars[i%len(chars)])
				i++
			}
		}
	}()
	return func() {
		once.Do(func() { close(done) })
	}
}
