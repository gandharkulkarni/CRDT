package lww

/* Properties are not exported */
type State struct {
	peer      string
	timestamp int64
	value     string
}
type LWWRegister struct {
	id    string
	state State
}

// ** Private methods **//
func (local *LWWRegister) setState(value string) {
	local.state = State{
		peer:      local.id,
		timestamp: local.state.timestamp + 1,
		value:     value,
	}
	// local.state.Timestamp += 1
	// local.state.Value = value
}

// ** Public methods **//
func (local *LWWRegister) GetValue() string {
	return local.state.value
}

func InitializeLWWRegister(id string, peer string, timestamp int, value string) *LWWRegister {
	return &LWWRegister{
		id: id,
		state: State{
			peer:      peer,
			timestamp: int64(timestamp),
			value:     value,
		},
	}
}
func (local *LWWRegister) UpdateLocalState(value string) {
	local.state.value = value
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
