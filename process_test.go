package coreclient

import (
	"os"
	"testing"

	gojson "encoding/json"

	"encoding/json"

	"github.com/datarhei/core-client-go/v16/api"
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

func BenchmarkDecodeProcessList(b *testing.B) {
	data, err := os.ReadFile("./fixtures/processList.json")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		processes := []api.Process{}
		err := json.Unmarshal(data, &processes)
		require.NoError(b, err)
	}
}

func BenchmarkDecodeProcessListNativeJSON(b *testing.B) {
	data, err := os.ReadFile("./fixtures/processList.json")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		processes := []api.Process{}
		err := gojson.Unmarshal(data, &processes)
		require.NoError(b, err)
	}
}

func BenchmarkEncodeProcessList(b *testing.B) {
	data, err := os.ReadFile("./fixtures/processList.json")
	require.NoError(b, err)

	processes := []api.Process{}
	err = gojson.Unmarshal(data, &processes)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(processes)
		require.NoError(b, err)
		require.NotEqual(b, 0, len(data))
	}
}

func BenchmarkEncodeProcessListNativeJSON(b *testing.B) {
	data, err := os.ReadFile("./fixtures/processList.json")
	require.NoError(b, err)

	processes := []api.Process{}
	err = gojson.Unmarshal(data, &processes)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		data, err := gojson.Marshal(processes)
		require.NoError(b, err)
		require.NotEqual(b, 0, len(data))
	}
}
