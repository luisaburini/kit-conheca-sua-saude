package controllers

import "fyne.io/fyne/v2/data/binding"

func NewStateManager(onStateChange func()) *StateManager {
	return &StateManager{
		stateListener: binding.NewDataListener(onStateChange),
	}
}

type StateManager struct {
	state         State
	stateListener binding.DataListener
}

type State uint8

const (
	Home State = 0
	//SentenceList State = 1
	Collection State = 1
)

func (sm *StateManager) GetState() State {
	return sm.state
}

func (sm *StateManager) SetState(state State) {
	sm.state = state
	sm.stateListener.DataChanged()
}
