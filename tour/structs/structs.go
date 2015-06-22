package main

import (
	"bytes"
	"fmt"
	"sync"
)

type Vertex struct {
	X int
	Y int
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{})     // {0, 0}
	fmt.Println(Vertex{X: 3}) // {3, 0}

	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)

	// pointer, will replace v.X
	q := &v
	q.X = 1e9
	fmt.Println(v)

	// non-pointer
	p := v
	p.X = 123
	fmt.Println(v)

	// 'new' function - allocates zero value and return pointer
	u := new(Vertex) // var u *U = new(Vertex)
	fmt.Println(u)
	u.X, u.Y = 11, 9
	fmt.Println(u)

	buf := new(SyncedBuffer) // type *SyncedBuffer
	var buf2 SyncedBuffer    // type SyncedBuffer
	fmt.Println(buf)
	fmt.Println(&buf2)
}
