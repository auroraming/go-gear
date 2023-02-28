package gcoroutine

import (
	"fmt"
	"testing"
)

func TestSyncMap(t *testing.T) {
	rMap := NewSyncMap[int, int]()
	rMap.LoadOrStore(1, 2)
	value, ok := rMap.Load(1)
	fmt.Println(value, ok)
	rMap.Delete(1)
	value, ok = rMap.Load(1)
	fmt.Println(value, ok)

}
