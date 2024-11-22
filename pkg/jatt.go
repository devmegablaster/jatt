package pkg

import (
	"fmt"
	"time"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/reader"
	"github.com/devmegablaster/jatt/pkg/renderer"
	"github.com/devmegablaster/jatt/pkg/styles"
)

type Jatt struct {
	reader   *reader.Reader
	renderer *renderer.Renderer
}

func NewJatt(cfg config.JattConfig) *Jatt {
	return &Jatt{
		reader:   reader.NewReader(cfg),
		renderer: renderer.NewRenderer(cfg),
	}
}

func (j *Jatt) Run() {
	timeStart := time.Now()

	files := j.reader.Read()
	renderedFiles := j.renderer.Render(files)
	fmt.Println(renderedFiles)

	timeEnd := time.Now()
	fmt.Println(styles.StatsStyle.Render(fmt.Sprintf("\n⚡Took %v", timeEnd.Sub(timeStart).Round(time.Millisecond))))
}
