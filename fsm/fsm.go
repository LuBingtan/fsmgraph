package fsm

/*****************************************  fsm interface  *****************************************/

type FSMState string
type FSMFunc func(interface{}) (interface{}, error)

type FSM interface {
	// create
	AddState(state FSMState)
	AddEvent(src, dst FSMState, f FSMFunc) error
	// read
	State() FSMState
	NextStates() []FSMState
	// delete
	RemoveState(state FSMState)
	// behavior
	Trigger(target FSMState) error
}