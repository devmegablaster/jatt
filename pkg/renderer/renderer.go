package renderer

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/a-h/templ"
	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/reader"
	"github.com/devmegablaster/jatt/pkg/styles"
	"github.com/devmegablaster/jatt/templates"
	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

type Renderer struct {
	cfg config.JattConfig
}

type RenderedFile struct {
	Name        string
	Html        []byte
	FrontMatter map[string]string
}

func NewRenderer(cfg config.JattConfig) *Renderer {
	return &Renderer{
		cfg: cfg,
	}
}

func (r *Renderer) Render(files []reader.File) []RenderedFile {
	renderedFiles := []RenderedFile{}

	for _, file := range files {

		html := r.ContentToHtml(file.Content)
		var component templ.Component

		switch file.FrontMatter["layout"] {
		case "home":
			content := templates.Home(r.cfg.HomeConfig)
			component = templates.Default(r.cfg.SiteConfig, r.cfg.NavConfig, content)
		case "listing":
			listing := templates.Listing(file.Listing)
			component = templates.Default(r.cfg.SiteConfig, r.cfg.NavConfig, listing)
		default:
			content := templates.Content(string(html))
			component = templates.Default(r.cfg.SiteConfig, r.cfg.NavConfig, content)
		}

		buf := bytes.Buffer{}
		component.Render(context.Background(), &buf)

		renderedFiles = append(renderedFiles, RenderedFile{
			Name:        file.Name,
			Html:        buf.Bytes(),
			FrontMatter: file.FrontMatter,
		})

		fmt.Println(styles.DebugStyle.Render("Rendered " + file.Name))
	}

	return renderedFiles
}

func (r *Renderer) ContentToHtml(content []byte) []byte {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			figure.Figure,
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error converting markdown to html"))
		os.Exit(1)
	}

	return buf.Bytes()
}
