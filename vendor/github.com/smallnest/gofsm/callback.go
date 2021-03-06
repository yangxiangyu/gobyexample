package fsm

// EventProcessor defines OnExit, Action and OnEnter actions.
type EventProcessor interface {
	// OnExit Action handles exiting a state
	OnExit(fromState interface{}, args []interface{})
	// Action is used to handle transitions
	Action(action interface{}, fromState interface{}, toState interface{}, args []interface{})
	// OnExit Action handles entering a state
	OnEnter(toState interface{}, args []interface{})
}

// DefaultDelegate is a default delegate.
// it splits processing of actions into three actions: OnExit, Action and OnEnter.
type DefaultDelegate struct {
	P EventProcessor
}

// HandleEvent implements Delegate interface and split HandleEvent into three actions.
func (dd *DefaultDelegate) HandleEvent(action interface{}, fromState interface{}, toState interface{}, args []interface{}) {
	if fromState != toState {
		dd.P.OnExit(fromState, args)
	}

	dd.P.Action(action, fromState, toState, args)

	if fromState != toState {
		dd.P.OnEnter(toState, args)
	}
}
