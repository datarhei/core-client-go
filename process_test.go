package coreclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseProcessID(t *testing.T) {
	tests := map[string]ProcessID{
		"foo":         NewProcessID("foo", ""),
		"foo@":        NewProcessID("foo", ""),
		"foo@bar":     NewProcessID("foo", "bar"),
		"foo@bar@bar": NewProcessID("foo@bar", "bar"),
	}

	for pid, id := range tests {
		ppid := ParseProcessID(pid)

		require.Equal(t, id, ppid, pid)
	}
}
