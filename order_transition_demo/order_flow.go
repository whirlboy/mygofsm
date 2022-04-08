package transition_demo

import (
	"fmt"
	"log"
)

type StatusFlow struct {
	ID                int64
	EventCount        int64
	StatusCore        StatusCore
	HistoryStatusCore []StatusCore
}

type EventProcessor interface {
	OnExit(fromState StatusCore, args []interface{})
	Action(action string, fromState StatusCore, toState StatusCore, args []interface{}) error
	OnActionFailure(action string, fromState StatusCore, toState StatusCore, args []interface{}, err error)
	OnEnter(toState StatusCore, args []interface{})
}

type StatusFlowEventProcessor struct{}

func (p *StatusFlowEventProcessor) OnExit(fromState StatusCore, args []interface{}) {
	t := args[0].(*StatusFlow)
	if t.StatusCore != fromState {
		panic(fmt.Errorf("订单 %v 的状态与期望的状态 %v 不一致，可能在状态机外被改变了", t, fromState))
	}
	log.Printf("订单 %d 初始状态 %+v", t.ID, fromState)
}

func (p *StatusFlowEventProcessor) Action(action string, fromState StatusCore, toState StatusCore, args []interface{}) error {
	t := args[0].(*StatusFlow)
	t.EventCount++

	switch action {
	case "demander_approve_script":
		fmt.Printf("正在执行%v动作\n", action)
	default: //其它action
	}

	return nil
}

func (p *StatusFlowEventProcessor) OnEnter(toState StatusCore, args []interface{}) {
	t := args[0].(*StatusFlow)
	t.StatusCore = toState
	// 记录状态机的中状态历史
	t.HistoryStatusCore = append(t.HistoryStatusCore, toState)

	log.Printf("订单 %d 状态改变为 %+v ", t.ID, toState)
}

func (p *StatusFlowEventProcessor) OnActionFailure(action string, fromState StatusCore, toState StatusCore, args []interface{}, err error) {
	t := args[0].(*StatusFlow)

	log.Printf("订单 %d 的状态从 %v to %v 改变失败， 原因: %v", t.ID, fromState, toState, err)
}
