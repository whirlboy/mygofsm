package transition_demo

type DefaultDelegate struct {
	P EventProcessor
}

func (d *DefaultDelegate) HandleEvent(action string, fromState StatusCore, toState StatusCore, args []interface{}) error {
	if fromState != toState {
		d.P.OnExit(fromState, args)
	}

	err := d.P.Action(action, fromState, toState, args)
	if err != nil {
		d.P.OnActionFailure(action, fromState, toState, args, err)
		return err
	}

	if fromState != toState {
		d.P.OnEnter(toState, args)
	}

	return nil
}
