package main

import (
	"fmt"
	"runtime"
	"time"

	"kaiyai.com/mymod/v2/goroutine"
)

// func sleep(s string) {
// 	var i int = 0
// 	for {
// 		i += 1i
// 		fmt.Println(i)
// 		fmt.Println(s)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	go goroutine.Sleep("Hello")
	go goroutine.Sleep("World")
	time.Sleep(1000 * time.Second)
}
