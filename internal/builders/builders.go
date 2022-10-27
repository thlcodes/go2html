package builders

import (
	"context"

	"github.com/dave/jennifer/jen"
	"golang.org/x/net/html"
)

type Builder interface {
	Build(ctx context.Context, doc *html.Node) (*jen.File, error)
}
