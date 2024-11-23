package writer

import (
	"fmt"
	"os"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/styles"
	"gopkg.in/yaml.v3"
)

func (w *Writer) InitJatt(dirname string) {
	w.CreateDir(dirname)

	jattConfig := config.NewConfigTemplate()

	yaml, err := yaml.Marshal(jattConfig)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error marshalling yaml"))
		os.Exit(1)
	}

	w.WriteFileInit(dirname+"/jatt.yaml", yaml)
	w.CreateDir(dirname + "/content")
	w.WriteFileInit(dirname+"/content/index.md", []byte("---\nlayout: home\n---\n"))
	w.WriteFileInit(dirname+"/content/about.md", []byte("---\nlayout: default\n---\n# About Me"))

	fmt.Println(styles.SuccessStyle.Render("\nJatt project initialized successfully"))
}

func (w *Writer) CreateDir(dirname string) {
	if err := os.Mkdir(dirname, 0755); err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error creating directory"))
		os.Exit(1)
	}

	fmt.Println(styles.DebugStyle.Render("Directory " + dirname + " created successfully"))
}

func (w *Writer) WriteFileInit(filename string, content []byte) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error creating file"))
		os.Exit(1)
	}

	defer file.Close()

	if _, err := file.Write(content); err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error writing to file"))
		os.Exit(1)
	}

	fmt.Println(styles.DebugStyle.Render("File " + filename + " created successfully"))
}
