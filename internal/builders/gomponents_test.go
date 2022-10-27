package builders_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"thl.codes/html2go/internal/builders"
)

func TestGomponentsBuilder(t *testing.T) {
	but := builders.NewGomponentsBuilder(zerolog.Nop())
	got, err := but.Build(context.TODO(), simpleTestDoc)
	require.NoError(t, err)
	require.NotNil(t, got)

	buf := bytes.NewBuffer(nil)
	require.NoError(t, got.Render(buf))

	fmt.Println(buf.String())
	require.Equal(t, `package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

var document = Doctype(
	HTML(
		Lang("en"),
		Body(
			g.El(
				"my-comp",
				Disabled(),
				g.Attr("custom", "bla"),
				DataAttr("bla", "bla"),
				Aria("bla", "bla"),
				g.Text("hi"),
			),
		),
	),
)
`, buf.String())
}

var simpleTestDoc = &html.Node{
	FirstChild: &html.Node{
		Type: html.DoctypeNode,
		Data: "html",
		NextSibling: &html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.Html,
			Attr: []html.Attribute{
				{Key: "lang", Val: "en"},
			},
			FirstChild: &html.Node{
				DataAtom: atom.Body,
				Type:     html.ElementNode,
				FirstChild: &html.Node{
					Type: html.ElementNode,
					Data: "my-comp",
					Attr: []html.Attribute{
						{Key: "disabled"},
						{Key: "custom", Val: "bla"},
						{Key: "data-bla", Val: "bla"},
						{Key: "aria-bla", Val: "bla"},
					},
					FirstChild: &html.Node{
						Type: html.TextNode,
						Data: "hi",
					},
				},
			},
		},
	},
	Type: html.DocumentNode,
}
