package loaders

import (
	"context"
	"io"

	"golang.org/x/net/html"
)

type Loader interface {
	Load(ctx context.Context, r io.ReadCloser) (*html.Node, error)
}
