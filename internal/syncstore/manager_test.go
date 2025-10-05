package syncstore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ximura/mnemosyne/internal/syncstore"
)

func TestManager_SaveAndLoad(t *testing.T) {
	mgr := syncstore.NewMemoryManager()

	// Initial state
	want := &syncstore.State{
		Synced: map[string]bool{
			"user:1": true,
			"user:2": false,
		},
	}

	// Save the state
	err := mgr.Save(want)
	require.NoError(t, err, "should save state without error")

	// Load it back
	got, err := mgr.Load()
	require.NoError(t, err, "should load state without error")
	require.NotNil(t, got, "loaded state should not be nil")

	require.Equal(t, want.Synced, got.Synced, "saved and loaded state must match")

	// Modify and re-save
	want.Synced["user:2"] = true
	err = mgr.Save(want)
	require.NoError(t, err)

	// Load again
	got, err = mgr.Load()
	require.NoError(t, err)
	require.Equal(t, want.Synced, got.Synced, "should reflect updated state")
}

func TestManager_EmptyFile(t *testing.T) {
	mgr := syncstore.NewMemoryManager()

	// No save yet → load should return empty state
	s, err := mgr.Load()
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Empty(t, s.Synced)
}
