package main

import (
	"fmt"
	"time"
)

func main() {
	//tickとafterはチャネルを送る
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
OuterLoop:
	for {
		select {
		case t := <-tick:
			fmt.Println("tick.", t)
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
			return

		//	二つのチャネル以外からきたものを扱う
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
