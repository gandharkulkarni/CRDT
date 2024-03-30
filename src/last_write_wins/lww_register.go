package lww

/* Properties are not exported */
type State struct {
	peer      string
	timestamp int64
	value     interface{}
}
type LWWRegister struct {
	id    string
	state State
}

// ** Private methods **//
func (local *LWWRegister) setState(value interface{}) {
	local.state = State{
		peer:      local.id,
		timestamp: local.state.timestamp + 1,
		value:     value,
	}
	local.state.timestamp += 1
	local.state.value = value

}

// ** Public methods **//
func (local *LWWRegister) GetValue() interface{} {
	return local.state.value
}
func InitializeLWWRegister(id string, state State) *LWWRegister {
	return &LWWRegister{
		id:    id,
		state: state,
	}
}
func (local *LWWRegister) Merge(state State) {
	if local.state.timestamp > state.timestamp {
		return
	}

	if local.state.timestamp == state.timestamp && local.state.peer > state.peer {
		return
	}

	local.state = state

}
