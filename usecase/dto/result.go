package dto

type State int

const (
	Success State = iota
	Failure
)

type Result[Output any, Fault any] struct {
	state State
	fault Fault
	value Output
}

func (result *Result[Output, Fault]) SetValue(value Output) {
	result.value = value
	result.state = Success
}

func (result *Result[Output, Fault]) Value() Output {
	return result.value
}

func (result *Result[Output, Fault]) SetFault(failure Fault) {
	result.fault = failure
	result.state = Failure
}

func (result *Result[Output, Fault]) Fault() Fault {
	return result.fault
}

func (result *Result[Output, Fault]) State() State {
	return result.state
}
