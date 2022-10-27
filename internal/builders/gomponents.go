package builders

import (
	"context"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

type GomponentsBuilder struct {
	log zerolog.Logger
}

var _ Builder = (*GomponentsBuilder)(nil)

func NewGomponentsBuilder(log zerolog.Logger) *GomponentsBuilder {
	return &GomponentsBuilder{
		log: log,
	}
}

var opts = jen.Options{
	Open:      "(",
	Close:     ")",
	Separator: ",",
	Multi:     true,
}

func (b *GomponentsBuilder) Build(ctx context.Context, doc *html.Node) (f *jen.File, err error) {
	f = jen.NewFile("main")
	f.ImportAlias("github.com/maragudk/gomponents", "g")
	f.ImportAlias("github.com/maragudk/gomponents/components", "c")
	f.ImportAlias("github.com/maragudk/gomponents/html", ".")

	g := "github.com/maragudk/gomponents"
	h := "github.com/maragudk/gomponents/html"

	var fn func(*html.Node) *jen.Statement
	fn = func(node *html.Node) *jen.Statement {
		var s *jen.Statement
		switch node.Type {
		case html.ElementNode:
			children := []jen.Code{}
			for _, attr := range node.Attr {
				var ele *jen.Statement
				var params []jen.Code
				if method, found := attributes[attr.Key]; found {
					ele = jen.Qual(h, method)
				} else if strings.HasPrefix(attr.Key, "data-") {
					ele = jen.Qual(h, "DataAttr")
					params = append(params, jen.Lit(attr.Key[5:]))
				} else if strings.HasPrefix(attr.Key, "aria-") {
					ele = jen.Qual(h, "Aria")
					params = append(params, jen.Lit(attr.Key[5:]))
				} else {
					ele = jen.Qual(g, "Attr")
					params = append(params, jen.Lit(attr.Key))
				}
				_, valueless := valuelessAttrs[attr.Key]
				if !valueless {
					params = append(params, jen.Lit(attr.Val))
				}
				ele.Call(params...)
				children = append(children, ele)
			}
			for n := node.FirstChild; n != nil; n = n.NextSibling {
				children = append(children, fn(n))
			}
			tag := node.DataAtom.String()
			if tag == "" {
				tag = node.Data
			}
			if method, found := elements[tag]; found {
				s = jen.Qual(h, method).Custom(opts, children...)
			} else {
				s = jen.Qual(g, "El").Custom(opts, append([]jen.Code{jen.Lit(tag)}, children...)...)
			}
		case html.TextNode:
			if strings.TrimSpace(node.Data) != "" {
				s = jen.Qual(g, "Text").Call(jen.Lit(node.Data))
			}
		case html.DoctypeNode:
			s = jen.Id("Doctype").Custom(opts, fn(node.NextSibling))
		case html.ErrorNode:
			b.log.Warn().Str("tag", node.Data).Msg("found invalid node, ignoring...")
		}
		return s
	}
	f.Var().Id("document").Op("=").Add(fn(doc.FirstChild))

	return
}

var elements = map[string]string{
	"a":          "A",
	"address":    "Address",
	"area":       "Area",
	"article":    "Article",
	"aside":      "Aside",
	"audio":      "Audio",
	"base":       "Base",
	"blockquote": "BlockQuote",
	"body":       "Body",
	"br":         "Br",
	"button":     "Button",
	"canvas":     "Canvas",
	"cite":       "Cite",
	"code":       "Code",
	"col":        "Col",
	"colgroup":   "ColGroup",
	"data":       "DataEl",
	"datalist":   "DataList",
	"details":    "Details",
	"dialog":     "Dialog",
	"div":        "Div",
	"dl":         "Dl",
	"embed":      "Embed",
	"form":       "FormEl",
	"fieldset":   "FieldSet",
	"figure":     "Figure",
	"footer":     "Footer",
	"head":       "Head",
	"header":     "Header",
	"hgroup":     "HGroup",
	"hr":         "Hr",
	"html":       "HTML",
	"iframe":     "IFrame",
	"img":        "Img",
	"input":      "Input",
	"label":      "Label",
	"legend":     "Legend",
	"li":         "Li",
	"link":       "Link",
	"main":       "Main",
	"menu":       "Menu",
	"meta":       "Meta",
	"meter":      "Meter",
	"nav":        "Nav",
	"noscript":   "NoScript",
	"object":     "Object",
	"ol":         "Ol",
	"optgroup":   "OptGroup",
	"option":     "Option",
	"p":          "P",
	"param":      "Param",
	"picture":    "Picture",
	"pre":        "Pre",
	"progress":   "Progress",
	"script":     "Script",
	"section":    "Section",
	"select":     "Select",
	"source":     "Source",
	"span":       "Span",
	"style":      "StyleEl",
	"summary":    "Summary",
	"svg":        "SVG",
	"table":      "Table",
	"tbody":      "TBody",
	"td":         "Td",
	"textarea":   "Textarea",
	"tfoot":      "TFoot",
	"th":         "Th",
	"thead":      "THead",
	"tr":         "Tr",
	"ul":         "Ul",
	"wbr":        "Wbr",
	"abbr":       "Abbr",
	"b":          "B",
	"caption":    "Caption",
	"dd":         "Dd",
	"del":        "Del",
	"dfn":        "Dfn",
	"dt":         "Dt",
	"em":         "Em",
	"figcaption": "FigCaption",
	"h1":         "H1",
	"h2":         "H2",
	"h3":         "H3",
	"h4":         "H4",
	"h5":         "H5",
	"h6":         "H6",
	"i":          "I",
	"ins":        "Ins",
	"kbd":        "Kbd",
	"mark":       "Mark",
	"q":          "Q",
	"s":          "S",
	"samp":       "Samp",
	"small":      "Small",
	"strong":     "Strong",
	"sub":        "Sub",
	"sup":        "Sup",
	"time":       "Time",
	"title":      "TitleEl",
	"u":          "U",
	"var":        "Var",
	"video":      "Video",
}

var attributes = map[string]string{
	"async":        "Async",
	"autofocus":    "AutoFocus",
	"autoplay":     "AutoPlay",
	"controls":     "Controls",
	"defer":        "Defer",
	"disabled":     "Disabled",
	"loop":         "Loop",
	"multiple":     "Multiple",
	"muted":        "Muted",
	"playsinline":  "PlaysInline",
	"readonly":     "ReadOnly",
	"required":     "Required",
	"selected":     "Selected",
	"accept":       "Accept",
	"action":       "Action",
	"alt":          "Alt",
	"as":           "As",
	"autocomplete": "AutoComplete",
	"charset":      "Charset",
	"class":        "Class",
	"cols":         "Cols",
	"content":      "Content",
	"for":          "For",
	"form":         "FormAttr",
	"height":       "Height",
	"href":         "Href",
	"id":           "ID",
	"lang":         "Lang",
	"loading":      "Loading",
	"max":          "Max",
	"maxlength":    "MaxLength",
	"method":       "Method",
	"min":          "Min",
	"minlength":    "MinLength",
	"name":         "Name",
	"pattern":      "Pattern",
	"placeholder":  "Placeholder",
	"poster":       "Poster",
	"preload":      "Preload",
	"rel":          "Rel",
	"role":         "Role",
	"rows":         "Rows",
	"src":          "Src",
	"srcset":       "SrcSet",
	"style":        "StyleAttr",
	"tabindex":     "TabIndex",
	"target":       "Target",
	"title":        "TitleAttr",
	"type":         "Type",
	"value":        "Value",
	"width":        "Width",
	"enctype":      "EncType",
}

var valuelessAttrs = map[string]struct{}{
	"async":       {},
	"autofocus":   {},
	"autoplay":    {},
	"controls":    {},
	"defer":       {},
	"disabled":    {},
	"loop":        {},
	"multiple":    {},
	"muted":       {},
	"playsinline": {},
	"readonly":    {},
	"required":    {},
	"selected":    {},
}
