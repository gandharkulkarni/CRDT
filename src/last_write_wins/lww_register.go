package lww

/* Properties are not exported */
type State struct {
	Peer      string
	Timestamp int64
	Value     string
}
type LWWRegister struct {
	Id    string
	State State
}

// ** Private methods **//
func (local *LWWRegister) setState(value string) {
	local.State = State{
		Peer:      local.Id,
		Timestamp: local.State.Timestamp + 1,
		Value:     value,
	}
	// local.state.Timestamp += 1
	// local.state.Value = value
}

// ** Public methods **//
func (local *LWWRegister) GetValue() string {
	return local.State.Value
}
func InitializeLWWRegister(id string, state State) *LWWRegister {
	return &LWWRegister{
		Id:    id,
		State: state,
	}
}
func (local *LWWRegister) Merge(state State) {
	if local.State.Timestamp > state.Timestamp {
		return
	}

	if local.State.Timestamp == state.Timestamp && local.State.Peer > state.Peer {
		return
	}

	local.State = state

}
