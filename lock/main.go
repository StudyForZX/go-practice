package lock

import (
	"fmt"
	"sync"
)

type TestS struct {
	Name  string
	Count int
}

func Lock(t *TestS, name string, l *sync.Mutex, wg *sync.WaitGroup) {

	l.Lock()

	defer l.Unlock()

	t.Count++
	t.Name = name

	res := fmt.Sprintf("testName: %s, Count is %d", t.Name, t.Count)

	fmt.Println(res)

	wg.Done()
}
