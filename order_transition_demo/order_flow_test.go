package order_transition_demo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFSM(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	statusFlow := &StatusFlow{
		ID:                 8888,
		StatusGroup:        StatusGroupA,
		HistoryStatusGroup: []StatusGroup{StatusGroupC},
	}
	fmt.Printf("老状态:%v\n", statusFlow)
	fsm := initFSM()
	err := fsm.Trigger(statusFlow.StatusGroup, "事件A", statusFlow)
	if err != nil {
		t.Errorf("trigger err: %v", err)
	}
	fmt.Printf("新状态:%v\n", statusFlow)

}
