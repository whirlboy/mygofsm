package order_transition_demo

func initFSM() *StateMachine {
	delegate := &DefaultDelegate{P: &StatusFlowEventProcessor{}}

	transitions := []Transition{
		{From: StatusGroupA, Event: "事件A", To: StatusGroupB, Action: "事件A"},
		{From: StatusGroupA, Event: "demander_reject_script", To: StatusGroupC, Action: "demander_reject_script"},
	}

	return NewStateMachine(delegate, transitions...)
}
