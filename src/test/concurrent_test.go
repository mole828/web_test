package concurrent_test

import (
	"testing"
	"time"
)

func concurrent_map_writes() {
	// fatal error: concurrent map writes
	data := map[int]int{}
	with_d := func(d int) {
		for i := 0; i < 100; i++ {
			data[i] = i + d
		}
	}
	go with_d(0)
	go with_d(1)
	go with_d(2)
	go with_d(3)
}

func concurrent_list_writes() *[]int {
	data := make([]int, 100)
	with_d := func(d int) {
		for i := 0; i < 100; i++ {
			data[i] = i + d
			time.Sleep(time.Microsecond)
		}
	}
	go with_d(0)
	go with_d(1)
	go with_d(2)
	go with_d(4)
	time.Sleep(time.Second)
	return &data
}

func TestConcurrent(t *testing.T) {
	if false {
		concurrent_map_writes()
	}
	t.Log(concurrent_list_writes())
}
