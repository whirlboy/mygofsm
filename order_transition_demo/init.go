package transition_demo

func initFSM() *StateMachine {
	delegate := &DefaultDelegate{P: &StatusFlowEventProcessor{}}

	transitions := []Transition{
		{From: s1, Event: "demander_approve_script", To: s2, Action: "demander_approve_script"},
		{From: s1, Event: "demander_reject_script", To: s3, Action: "demander_reject_script"},
	}

	return NewStateMachine(delegate, transitions...)
}
