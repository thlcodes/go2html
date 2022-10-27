package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog"
	"thl.codes/html2go/internal/builders"
	"thl.codes/html2go/internal/loaders"
)

type MainCommand struct {
	log     zerolog.Logger
	loader  loaders.Loader
	builder builders.Builder
	logOnly bool
	dir     string
}

var _ Command = (*MainCommand)(nil)

func NewMainCommand(log zerolog.Logger, dir string, logOnly bool) *MainCommand {
	cmd := &MainCommand{
		log:     log,
		loader:  loaders.NewHtmlLoader(log),
		builder: builders.NewGomponentsBuilder(log),
		logOnly: logOnly,
		dir:     dir,
	}
	return cmd
}

func (cmd *MainCommand) Run(ctx context.Context, files ...string) error {
	if len(files) == 0 {
		return fmt.Errorf("not input files")
	}
	readers := map[string]io.ReadCloser{}
	for _, file := range files {
		r, err := load(file)
		if err != nil {
			return err
		}
		readers[file] = r
	}
	for f, r := range readers {
		node, err := cmd.loader.Load(ctx, r)
		if err != nil {
			return fmt.Errorf("could not load html from %s: %w", f, err)
		}
		out, err := cmd.builder.Build(ctx, node)
		if err != nil {
			return fmt.Errorf("could not generae code from %s: %w", f, err)
		}
		if cmd.logOnly {
			buf := bytes.NewBuffer(nil)
			if err := out.Render(buf); err != nil {
				return fmt.Errorf("could not render %s to code: %w", f, err)
			}
			fmt.Println(buf.String())
		} else {
			target := f + ".go"
			if cmd.dir != "." {
				target = path.Join(cmd.dir, path.Base(target))
			}
			if err := out.Save(target); err != nil {
				return fmt.Errorf("could not write to %s: %w", target, err)
			}
		}
	}
	return nil
}

func load(file string) (io.ReadCloser, error) {
	if strings.HasPrefix(file, "https://") || strings.HasPrefix(file, "http://") {
		return loadHttp(file)
	} else {
		return loadFile(file)
	}
}

func loadFile(path string) (r io.ReadCloser, err error) {
	if info, err := os.Stat(path); err != nil || info.IsDir() {
		return nil, fmt.Errorf("%s does not exist or is a dir", path)
	}
	r, err = os.Open(path)
	if err != nil {
		err = fmt.Errorf("could not open %s for reading: %w", path, err)
	}
	return
}

func loadHttp(url string) (r io.ReadCloser, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch %s", url)
	}
	r = resp.Body
	return
}
