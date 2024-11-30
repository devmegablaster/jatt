package pkg

import (
	"fmt"
	"time"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/reader"
	"github.com/devmegablaster/jatt/pkg/renderer"
	"github.com/devmegablaster/jatt/pkg/rss"
	"github.com/devmegablaster/jatt/pkg/styles"
	"github.com/devmegablaster/jatt/pkg/writer"
)

type Jatt struct {
	reader   *reader.Reader
	rss      *rss.Rss
	renderer *renderer.Renderer
	writer   *writer.Writer
}

func NewJatt(cfg config.JattConfig) *Jatt {
	return &Jatt{
		reader:   reader.NewReader(cfg),
		renderer: renderer.NewRenderer(cfg),
		writer:   writer.NewWriter(cfg),
		rss:      rss.New(cfg),
	}
}

func (j *Jatt) Run() {
	timeStart := time.Now()

	files := j.reader.Read()
	feed := j.rss.GenerateFeed(files)
	renderedFiles := j.renderer.Render(files)
	j.writer.WriteFiles(renderedFiles)
	j.writer.WriteRSSFeed(feed)
	j.writer.CopyStatic()

	timeEnd := time.Now()
	fmt.Println(styles.StatsStyle.Render(fmt.Sprintf("\n⚡Took %v", timeEnd.Sub(timeStart).Round(time.Millisecond))))
}
