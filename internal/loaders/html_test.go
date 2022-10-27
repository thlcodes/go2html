package loaders_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	_ "embed"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"

	"thl.codes/html2go/internal/loaders"
)

//go:embed testdata/simple.html
var simpleTestFile string

func TestHtmlLoader(t *testing.T) {
	lut := loaders.NewHtmlLoader(zerolog.Nop())
	doc, err := lut.Load(context.TODO(), io.NopCloser(bytes.NewBufferString(simpleTestFile)))
	require.NoError(t, err)
	require.NotNil(t, doc)
	require.Equal(t, html.DocumentNode, doc.Type)
}
