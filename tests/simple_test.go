package tests

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"thl.codes/html2go/internal/commands"
)

type logEntry struct {
	Message string `json:"message"`
}

func TestSimple(t *testing.T) {
	logbuf := bytes.NewBuffer(nil)
	cmd := commands.NewMainCommand(zerolog.New(logbuf), ".", true)
	require.NoError(t, cmd.Run(context.TODO(), "./testdata/simple.html"))
	var entry logEntry
	require.NoError(t, json.NewDecoder(logbuf).Decode(&entry))
	require.NotEmpty(t, entry.Message)
	p := path.Join(t.TempDir(), "simple.html.go")
	require.NoError(t, os.WriteFile(p, []byte(entry.Message), 0))
	fmt.Println(entry.Message)
	defer os.Remove(p)
}
