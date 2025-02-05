package lock

import (
	"sync"
	"testing"
)

func TestLock(t *testing.T) {

	ts := &TestS{Count: 0}

	var wg sync.WaitGroup
	var l sync.Mutex

	wg.Add(11)

	go Lock(ts, "name1", &l, &wg)
	go Lock(ts, "name2", &l, &wg)
	go Lock(ts, "name3", &l, &wg)
	go Lock(ts, "name4", &l, &wg)
	go Lock(ts, "name5", &l, &wg)
	go Lock(ts, "name6", &l, &wg)
	go Lock(ts, "name7", &l, &wg)
	go Lock(ts, "name8", &l, &wg)
	go Lock(ts, "name9", &l, &wg)
	go Lock(ts, "name10", &l, &wg)
	go Lock(ts, "name11", &l, &wg)

	wg.Wait()

	if ts.Count != 11 {
		t.Errorf("Excepted 11, got %d", ts.Count)
	}
}
