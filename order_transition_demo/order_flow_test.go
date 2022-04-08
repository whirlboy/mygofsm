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
		ID:                88888888,
		StatusCore:        s1,
		HistoryStatusCore: []StatusCore{s1},
	}
	fmt.Printf("老状态:%v\n", statusFlow)
	fsm := initFSM()
	err := fsm.Trigger(statusFlow.StatusCore, "demander_approve_script", statusFlow)
	if err != nil {
		t.Errorf("trigger err: %v", err)
	}
	fmt.Printf("新状态:%v\n", statusFlow)

}
