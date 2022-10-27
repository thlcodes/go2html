package loaders

import (
	"context"
	"fmt"
	"io"

	"github.com/rs/zerolog"

	"golang.org/x/net/html"
)

type HtmlLoader struct {
	log zerolog.Logger
}

var _ Loader = (*HtmlLoader)(nil)

func NewHtmlLoader(log zerolog.Logger) *HtmlLoader {
	return &HtmlLoader{
		log: log,
	}
}

func (*HtmlLoader) Load(ctx context.Context, r io.ReadCloser) (*html.Node, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("could not parse document: %w", err)
	}
	return node, nil
}
