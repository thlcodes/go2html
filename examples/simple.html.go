package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

var simpleDocument = Doctype(
	HTML(
		Lang("en"),
		Head(
			g.Text("\n    "),
			Meta(
				Charset("UTF-8"),
			),
			g.Text("\n    "),
			Meta(
				g.Attr("http-equiv", "X-UA-Compatible"),
				Content("IE=edge"),
			),
			g.Text("\n    "),
			Meta(
				Name("viewport"),
				Content("width=device-width, initial-scale=1.0"),
			),
			g.Text("\n    "),
			TitleEl(
				g.Text("Document"),
			),
			g.Text("\n"),
		),
		g.Text("\n\n"),
		Body(
			g.Text("\n    Bodi\n    "),
			g.El(
				"my-comp",
			),
			g.Text("\n\n\n"),
		),
	),
)
