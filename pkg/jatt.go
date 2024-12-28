package pkg

import (
	"fmt"
	"time"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/reader"
	"github.com/devmegablaster/jatt/pkg/renderer"
	"github.com/devmegablaster/jatt/pkg/rss"
	"github.com/devmegablaster/jatt/pkg/sitemap"
	"github.com/devmegablaster/jatt/pkg/styles"
	"github.com/devmegablaster/jatt/pkg/writer"
)

type Jatt struct {
	reader   *reader.Reader
	rss      *rss.RssSvc
	renderer *renderer.Renderer
	writer   *writer.Writer
	sitemap  *sitemap.SitemapSvc
}

func NewJatt(cfg config.JattConfig) *Jatt {
	return &Jatt{
		reader:   reader.NewReader(cfg),
		renderer: renderer.NewRenderer(cfg),
		writer:   writer.NewWriter(cfg),
		rss:      rss.New(cfg),
		sitemap:  sitemap.New(cfg),
	}
}

func (j *Jatt) Run() {
	timeStart := time.Now()

	// Read and render files
	files := j.reader.Read()
	renderedFiles := j.renderer.Render(files)

	// Write files
	j.writer.WriteFiles(renderedFiles)
	// Copy static files
	j.writer.CopyStatic()

	// Generate and write RSS feed
	feed := j.rss.GenerateFeed(files, renderedFiles)
	j.writer.WriteFileWithName("rss.xml", feed)

	// Generate and write sitemap
	sitemap := j.sitemap.GenerateSitemap(files)
	j.writer.WriteFileWithName("sitemap.xml", sitemap)

	// Generate and write robots.txt
	robots := j.sitemap.GenerateRobots()
	j.writer.WriteFileWithName("robots.txt", robots)

	timeEnd := time.Now()
	fmt.Println(styles.StatsStyle.Render(fmt.Sprintf("\n⚡Took %v", timeEnd.Sub(timeStart).Round(time.Millisecond))))
}
