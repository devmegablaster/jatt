package writer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/renderer"
	"github.com/devmegablaster/jatt/pkg/styles"
)

type Writer struct {
	cfg config.JattConfig
}

func NewWriter(cfg config.JattConfig) *Writer {
	return &Writer{
		cfg: cfg,
	}
}

func (w *Writer) WriteFiles(wg *sync.WaitGroup, files []renderer.RenderedFile) {
	defer wg.Done()

	w.RemoveOldOutput()
	w.CreateOutputDir()

	for _, file := range files {
		if err := w.WriteFile(file); err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error writing " + file.Name))
			os.Exit(1)
		}

		fmt.Println(styles.SuccessStyle.Render("Wrote " + file.Name))
	}
}

func (w *Writer) RemoveOldOutput() {
	outputDir := w.cfg.SiteConfig.OutputDir
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		return
	}

	err := os.RemoveAll(outputDir)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error removing old output"))
		os.Exit(1)
	}
}

func (w *Writer) CreateOutputDir() {
	outputDir := w.cfg.SiteConfig.OutputDir
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error creating output directory"))
		os.Exit(1)
	}
}

func (w *Writer) WriteFile(file renderer.RenderedFile) error {
	outputDir := w.cfg.SiteConfig.OutputDir
	f, err := os.Create(outputDir + "/" + file.Name + ".html")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(outputDir+"/"+strings.Split(file.Name, "/")[0], 0755)
			f, err = os.Create(outputDir + "/" + file.Name + ".html")
			if err != nil {
				return err
			}
		}
	}

	defer f.Close()

	_, err = f.Write(file.Html)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) CopyStatic(wg *sync.WaitGroup) error {
	defer wg.Done()

	staticDir, err := os.ReadDir("static")
	if err != nil {
		return err
	}

	for _, file := range staticDir {
		if file.IsDir() {
			continue
		}

		f, err := os.Open("static/" + file.Name())
		if err != nil {
			return err
		}
		defer f.Close()

		buf := bytes.Buffer{}
		io.Copy(&buf, f)

		os.WriteFile(w.cfg.SiteConfig.OutputDir+"/"+file.Name(), buf.Bytes(), 0644)
	}

	return nil
}

func (w *Writer) WriteRSSFeed(feed []byte) {
	outputDir := w.cfg.SiteConfig.OutputDir
	err := os.WriteFile(outputDir+"/rss.xml", feed, 0644)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error writing RSS feed"))
		os.Exit(1)
	}

	fmt.Println(styles.SuccessStyle.Render("Wrote RSS feed"))
}
