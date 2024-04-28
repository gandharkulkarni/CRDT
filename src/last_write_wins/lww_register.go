package lww

import "fmt"

/* Properties are not exported */
type State struct {
	peer      string
	peerId    int64
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
		peerId:    1,
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

func (local *LWWRegister) GetTimestamp() int64 {
	return local.state.timestamp
}

func (local *LWWRegister) GetPeerId() int64 {
	return local.state.peerId
}

func (local *LWWRegister) PopulatePeerState(peer string, peerId int64, timestamp int, value string) State {
	return State{
		peer:      peer,
		timestamp: int64(timestamp),
		value:     value,
	}
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
	local.state.timestamp++
}
func (local *LWWRegister) Merge(state State) {
	fmt.Println("Peer state", state)
	if local.state.timestamp > state.timestamp {
		return
	}

	if local.state.timestamp == state.timestamp && local.state.peerId > state.peerId {
		return
	}

	local.state = state

}
