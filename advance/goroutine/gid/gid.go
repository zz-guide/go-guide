package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

/**
获取goroutine ID
参考链接：https://github.com/kortschak/goroutine/blob/master/gid.go
*/
func main() {
	TestID()
}

func TestID() {
	var got int64
	want := goid()
	if got != want {
		fmt.Printf("unexpected id for main goroutine: got:%d want:%d\n", got, want)
	}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			got = 0
			want := goid()
			if got != want {
				fmt.Printf("unexpected id for goroutine number %d: got:%d want:%d\n", i, got, want)
			}
		}()
	}
	wg.Wait()
}

// goid returns the goroutine ID extracted from a stack trace.
func goid() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// Only for debug, never use it in production
func RoutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	// if error, just return 0
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}
