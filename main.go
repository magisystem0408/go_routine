package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	//lockをしてからmux
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	//二つのgo routineから一つのmapを書き換えようとするとエラーが出る
	//c := make(map[string]int)
	c := Counter{v: make(map[string]int)}

	//go routine
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()

	time.Sleep(time.Second * 1)
	fmt.Println(c, c.Value("key"))
}
