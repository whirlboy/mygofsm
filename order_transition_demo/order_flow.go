package order_transition_demo

import (
	"fmt"
	"log"
)

type StatusFlow struct {
	ID                 int64
	EventCount         int64
	StatusGroup        StatusGroup
	HistoryStatusGroup []StatusGroup
}

type EventProcessor interface {
	OnExit(fromState StatusGroup, args []interface{})
	Action(action string, fromState StatusGroup, toState StatusGroup, args []interface{}) error
	OnActionFailure(action string, fromState StatusGroup, toState StatusGroup, args []interface{}, err error)
	OnEnter(toState StatusGroup, args []interface{})
}

type StatusFlowEventProcessor struct{}

func (p *StatusFlowEventProcessor) OnExit(fromState StatusGroup, args []interface{}) {
	t := args[0].(*StatusFlow)
	if t.StatusGroup != fromState {
		panic(fmt.Errorf("订单 %v 的状态与期望的状态 %v 不一致，可能在状态机外被改变了", t, fromState))
	}
	log.Printf("订单 %d 初始状态 %+v", t.ID, fromState)
}

func (p *StatusFlowEventProcessor) Action(action string, fromState StatusGroup, toState StatusGroup, args []interface{}) error {
	t := args[0].(*StatusFlow)
	t.EventCount++

	switch action {
	case "demander_approve_script":
		fmt.Printf("正在执行%v事件\n", action)
	default: //其它action
	}

	return nil
}

func (p *StatusFlowEventProcessor) OnEnter(toState StatusGroup, args []interface{}) {
	t := args[0].(*StatusFlow)
	t.StatusGroup = toState
	// 记录状态机的中状态历史
	t.HistoryStatusGroup = append(t.HistoryStatusGroup, toState)

	log.Printf("订单 %d 状态改变为 %+v ", t.ID, toState)
}

func (p *StatusFlowEventProcessor) OnActionFailure(action string, fromState StatusGroup, toState StatusGroup, args []interface{}, err error) {
	t := args[0].(*StatusFlow)

	log.Printf("订单 %d 的状态从 %v to %v 改变失败， 原因: %v", t.ID, fromState, toState, err)
}
