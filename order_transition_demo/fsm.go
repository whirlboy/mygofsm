package order_transition_demo

import (
	"fmt"
)

type StatusGroup struct {
	OrderStatus        int16
	OrderAuditStatus   int16
	OrderPaymentStatus int16
}

type Transition struct {
	From   StatusGroup
	Event  string
	To     StatusGroup
	Action string
}

type Delegate interface {
	HandleEvent(action string, fromState StatusGroup, toState StatusGroup, args []interface{}) error
}

type StateMachine struct {
	delegate    Delegate
	transitions []Transition
}

type Error interface {
	error
	BadEvent() string
	CurrentState() string
}

type smError struct {
	badEvent     string
	currentState StatusGroup
}

func (e smError) Error() string {
	return fmt.Sprintf("state machine error: cannot find transition for event [%v] when in state [%v]\n", e.badEvent, e.currentState)
}

func (e smError) BadEvent() string {
	return e.badEvent
}

func (e smError) CurrentState() StatusGroup {
	return e.currentState
}

func NewStateMachine(delegate Delegate, transitions ...Transition) *StateMachine {
	return &StateMachine{delegate: delegate, transitions: transitions}
}

func (m *StateMachine) Trigger(currentState StatusGroup, event string, args ...interface{}) error {
	currentTransition := m.findTransMatching(currentState, event)
	if currentTransition == nil {
		return smError{event, currentState}
	}

	var err error
	if currentTransition.Action != "" {
		err = m.delegate.HandleEvent(currentTransition.Action, currentState, currentTransition.To, args)
	}
	return err
}

func (m *StateMachine) findTransMatching(fromState StatusGroup, event string) *Transition {
	for _, v := range m.transitions {
		if v.From == fromState && v.Event == event {
			return &v
		}
	}
	return nil
}
