package syncstore

import (
	"encoding/json"
	"sync"
)

type State struct {
	synced sync.Map
}

func NewState() *State {
	return &State{}
}

func (s *State) IsSynced(id string) bool {
	v, ok := s.synced.Load(id)
	if !ok {
		return false
	}
	b, _ := v.(bool)
	return b
}

func (s *State) MarkSynced(id string) {
	s.synced.Store(id, true)
}

func (s *State) MarshalJSON() ([]byte, error) {
	// Convert to plain map for serialization
	return json.Marshal(struct {
		Synced map[string]bool `json:"synced"`
	}{
		Synced: s.data(),
	})
}

func (s *State) UnmarshalJSON(data []byte) error {
	aux := struct {
		Synced map[string]bool `json:"synced"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Rebuild sync.Map
	for k, v := range aux.Synced {
		s.synced.Store(k, v)
	}
	return nil
}

func (s *State) data() map[string]bool {
	m := make(map[string]bool)
	s.synced.Range(func(k, v any) bool {
		key, _ := k.(string)
		val, _ := v.(bool)
		m[key] = val
		return true
	})
	return m
}
