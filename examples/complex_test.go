package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestComplex(t *testing.T) {
	want, err := os.ReadFile("./complex.html")
	require.NoError(t, err)
	wantHtml, err := html.Parse(bytes.NewBuffer(want))
	require.NoError(t, err)
	buf := bytes.NewBuffer(nil)
	require.NoError(t, complexDocument.Render(buf))
	fmt.Println(buf.String())
	gotHtml, err := html.Parse(buf)
	require.NoError(t, err)
	require.Equal(t, wantHtml, gotHtml)
}
