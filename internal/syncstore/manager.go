package syncstore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Manager struct {
	newReader func() (io.ReadCloser, error)
	newWriter func() (io.WriteCloser, error)
}

func NewManager(newReader func() (io.ReadCloser, error),
	newWriter func() (io.WriteCloser, error),
) *Manager {
	return &Manager{newReader: newReader, newWriter: newWriter}
}

func (m *Manager) Load() (*State, error) {
	reader, err := m.newReader()
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %w", err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read state: %w", err)
	}

	s := NewState()
	if len(data) == 0 {
		return s, nil
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("failed to unmarshal state: %w", err)
	}

	return s, nil
}

func (m *Manager) Save(s *State) error {
	writer, err := m.newWriter()
	if err != nil {
		return fmt.Errorf("failed to create writer: %w", err)
	}
	defer writer.Close()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marhal state: %w", err)
	}

	_, err = writer.Write(data)
	return err
}

func NewFileManager() *Manager {
	path := filepath.Join(os.TempDir(), "mnemo_state.json")

	newReader := func() (io.ReadCloser, error) {
		return os.Open(path)
	}

	newWriter := func() (io.WriteCloser, error) {
		return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	}
	return NewManager(newReader, newWriter)
}

func NewMemoryManager() *Manager {
	store := bytes.Buffer{}

	newReader := func() (io.ReadCloser, error) {
		r := bytes.NewReader(store.Bytes())
		return io.NopCloser(r), nil
	}

	newWriter := func() (io.WriteCloser, error) {
		// Reset buffer each time we save
		store.Reset()
		return nopWriteCloser{&store}, nil
	}

	return NewManager(newReader, newWriter)
}

type nopWriteCloser struct{ io.Writer }

func (nopWriteCloser) Close() error { return nil }
