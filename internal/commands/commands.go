package commands

import "context"

type Command interface {
	Run(ct context.Context, args ...string) error
}
